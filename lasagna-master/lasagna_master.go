package main

import "fmt"

// TODO: define the 'PreparationTime()' function
/*## 1. Estimate the preparation time

For the next lasagna that you will prepare, you want to make sure you have enough time reserved so you can enjoy the cooking.
You already planned which layers your lasagna will have.
Now you want to estimate how long the preparation will take based on that.

Implement a function `PreparationTime` that accepts a slice of layers as a `[]string` and the average preparation time per layer in minutes as an `int`.
The function should return the estimate for the total preparation time based on the number of layers as an `int`.
Go has no default values for functions.
If the average preparation time is passed as `0` (the default initial value for an `int`), then the default value of `2` should be used.

```go
layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
PreparationTime(layers, 3)
// => 18
PreparationTime(layers, 0)
// => 12*/
func PreparationTime(slice_of_layers []string, average_preparation_time_in_minutes int) int {
	var slice_length, preparation_time int = len(slice_of_layers), 0
	if average_preparation_time_in_minutes == 0 {
		average_preparation_time_in_minutes = 2
	}
	preparation_time = slice_length * average_preparation_time_in_minutes
	return preparation_time

}

/*## 2. Compute the amounts of noodles and sauce needed

Besides reserving the time, you also want to make sure you have enough sauce and noodles to cook the lasagna of your dreams.
For each noodle layer in your lasagna, you will need 50 grams of noodles.
For each sauce layer in your lasagna, you will need 0.2 liters of sauce.

Define the function `Quantities` that takes a slice of layers as parameter as a `[]string`.
The function will then determine the quantity of noodles and sauce needed to make your meal.
The result should be returned as two values of `noodles` as an `int` and `sauce` as a `float64`.

```go
Quantities([]string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"})
// => 100, 0.4*/
// TODO: define the 'Quantities()' function
func Quantities(slice_of_layers []string) (noodles int, sauce float64) {
	var slice_length int = len(slice_of_layers)
	var counter int = 0
	for counter = 0; counter < slice_length; counter++ {
		if slice_of_layers[counter] == "noodles" {
			noodles += 50
		}
		if slice_of_layers[counter] == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce

}

// TODO: define the 'AddSecretIngredient()' function
/*
## 3. Add the secret ingredient

A while ago you visited a friend and ate lasagna there.
It was amazing and had something special to it.
The friend sent you the list of ingredients and told you the last item on the list is the "secret ingredient" that made the meal so special.
Now you want to add that secret ingredient to your recipe as well.

Write a function `AddSecretIngredient` that accepts two slices of ingredients of type `[]string` as parameters.
The first parameter is the list your friend sent you, the second is the ingredient list of your own recipe.
The last element in your ingredient list is always `"?"`. The function should replace it with the last item from your friends list.
**Note:** `AddSecretIngredient` does not return anything - you should modify the list of your ingredients directly. The list with your friend's ingredients should **not** be modified. Also, since `slice` is passed into a function as pointers, changes to the two `[]string` arguments passed into `AddSecretIngredient` will be modified directly.

```go
friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
myList := []string{"noodles", "meat", "sauce", "mozzarella","?"}

AddSecretIngredient(friendsList, myList)
// myList => []string{"noodles", "meat", "sauce", "mozzarella", "kampot pepper"}
```*/
func AddSecretIngredient(myList []string, friends_list []string) []string {
	var friends_list_length, myList_length int = len(friends_list) - 1, len(myList) - 1
	myList[(myList_length)] = friends_list[(friends_list_length)]

	return myList
}

// TODO: define the 'ScaleRecipe()' function
/*## 4. Scale the recipe

The amounts listed in your cookbook only yield enough lasagna for two portions.
Since you want to cook for more people next time, you want to calculate the amounts for different numbers of portions.

Implement a function `ScaleRecipe` that takes two parameters.

- A slice of `float64` amounts needed for 2 portions.
- The number of portions you want to cook.

The function should return a slice of `float64` of the amounts needed for the desired number of portions.
You want to keep the original recipe though.
This means the `quantities` argument should not be modified in this function.

```go
quantities := []float64{ 1.2, 3.6, 10.5 }
scaledQuantities := ScaleRecipe(quantities, 4)
// => []float64{ 2.4, 7.2, 21 }
*/
func ScaleRecipe(amount_needed_for_2_portions []float64, number_of_portions int) []float64 {
	var loop_counter int = 0
	var new_slice_for_2_portions_of_lasagna []float64 = make([]float64, len(amount_needed_for_2_portions))
	copy(new_slice_for_2_portions_of_lasagna, amount_needed_for_2_portions)
	var product float64 = float64(number_of_portions)
	for loop_counter = 0; loop_counter < len(new_slice_for_2_portions_of_lasagna); loop_counter++ {
		new_slice_for_2_portions_of_lasagna[loop_counter] *= (product / 2)
	}
	return new_slice_for_2_portions_of_lasagna
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
/*=== RUN   TestAddSecretIngredient/Adds_secret_ingredient

    lasagna_master_test.go:158: addSecretIngredient([sauce noodles béchamel marjoram], [sauce noodles meat tomatoes ?]) = [sauce noodles meat tomatoes ?]
	 want [sauce noodles meat tomatoes marjoram]

--- FAIL: TestAddSecretIngredient/Adds_secret_ingredient (0.00s)*/
func main() {
	/*var friends_list, myList []string = []string{"sauce", "noodles", "béchamel", "marjoram"}, []string{"sauce", "noodles", "meat", "tomatoes", "?"}
	AddSecretIngredient(myList, friends_list)
	fmt.Println(myList)*/
	//input:    []float64{0.5, 250, 150, 3, 0.5}, portions: 6,
	var quantities []float64 = []float64{0.5, 250, 150, 3, 0.5}
	fmt.Println(ScaleRecipe(quantities, 6))

}
