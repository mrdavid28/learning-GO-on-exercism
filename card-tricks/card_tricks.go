package main

import "fmt"

//package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	cards := []int{2, 6, 9}
	return cards
}

// GetItem retrieves an item from a slice at given position.
func GetItem(slice []int, index int) int {
	if index < 0 || index > len(slice)-1 {
		return -1
	} else {
		return slice[index]
	}
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index < 0 || index > len(slice)-1 {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	if len(values) == 0 {
		return slice
	} else {
		slice = append(values, slice...)
		return slice
	}
}

/*## 5. Remove a card from the stack

Remove the card at position `index` from the stack and return the stack.
Note that this may modify the input slice which is ok.

```go
cards := RemoveItem([]int{3, 2, 6, 4, 8}, 2)
fmt.Println(cards)
// Output: [3 2 4 8]
```

If the index is out of bounds (ie. if it is negative or after the end of the stack), we want to leave the stack unchanged:

```go
cards := RemoveItem([]int{3, 2, 6, 4, 8}, 11)
fmt.Println(cards)
// Output: [3 2 6 4 8]*/
// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index < 0 || index > len(slice)-1 {
		return slice
	}
	slice = append(slice[:index], slice[index+1:]...)
	return slice
}

func main() {
	/*fmt.Println(GetItem([]int{2, 4, 5, 6, 13, 19}, 6))
	fmt.Println(SetItem([]int{0, 1, 2, 3}, -1, 15))
	fmt.Println(PrependItems([]int{3, 2, 6, 4}, 6, 7, 8, 9, 19)) //prepending items at top of the slice in that case there are 6,7,8,9,19
	*/
	fmt.Println(RemoveItem([]int{3, 2, 6, 4, 8}, 16))
}
