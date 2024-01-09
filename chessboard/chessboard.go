package chessboard

import "fmt"

// Declare a type named File which stores if a square is occupied by a piece -
// this will be a slice of bools

// Declare a type named Chessboard which contains a map of eight Files,
// accessed with keys from "A" to "H"

type File [8]bool
type Chessboard map[string]File // that Chessboard map has keys of type string and values of type bool because slice File is bool type0

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
/*## 1. Given a Chessboard and a File, count how many squares are occupied
Implement the `CountInFile(board Chessboard, file string) int` function.
It should count the total number of occupied squares by ranging over a map. Return an integer.
Return a count of zero (`0`) if the given file cannot be found in the map.

```go
CountInFile(board, "A")
// => 3*/
func CountInFile(cb Chessboard, file string) int {
	var value_of_occurences int = 0
	for _, isOccupied := range cb[file] { //it's same type because for key file of type string we have value of type bool so there is ok
		if isOccupied { // if is occupied is true then increment value_of_occurences kind of easy
			value_of_occurences++
		}
	}
	return value_of_occurences
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 0 || rank > 8 {
		return 0
	}
	var count int = 0
	for file := range cb {
		if cb[file][rank-1] {
			count++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var count int = 0
	for _, t := range cb {
		count += len(t)
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var count_occupied int = 0
	for t := range cb {
		for _, isOccupied := range cb[t] {
			if isOccupied {
				count_occupied++
			}
		}
	}
	return count_occupied
}

func main() {
	fmt.Println("Cos takiego")
}
