// package birdwatcher
package main

import "fmt"

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var sum, loop_counter int = 0, 0

	for loop_counter = 0; loop_counter < len(birdsPerDay); loop_counter++ {
		sum = sum + birdsPerDay[loop_counter]
	}
	return len(birdsPerDay)
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week. One week has 7 seven days so we need to include that in our solution Bird count is for individual days
// That seems to be even more complicated because solution need to be for specific week not for number of weeks for example
// we need to count birds in week two but we will be not counting all birds like I had first thought
func BirdsInWeek(birdsPerDay []int, week int) int {
	var slice_size_counter, sum = 0, 0

	if week == 1 {
		slice_size_counter = 0
	} else {
		slice_size_counter = (week * 7) - 7
	}
	if slice_size_counter == 0 {
		for slice_size_counter = slice_size_counter; slice_size_counter < 7; slice_size_counter++ {
			sum = sum + birdsPerDay[slice_size_counter]
		}
	} else {
		var normal_counter int = (slice_size_counter * week)

		for slice_size_counter = slice_size_counter; slice_size_counter < normal_counter; slice_size_counter++ {

			sum = sum + birdsPerDay[slice_size_counter]
		}

	}
	return sum
}

/*
	You realized that all the time you were trying to keep track of the birds, there was one bird that was hiding in a far corner of the garden.

You figured out that this bird always spent every second day in your garden.

You do not know exactly where it was in between those days but definitely not in your garden.

Your bird watcher intuition also tells you that the bird was in your garden on the first day that you tracked in your list.

Given this new information, write a function FixBirdCountLog that takes a slice of birds counted per day as an argument and returns the slice after correcting the counting mistake.

birdsPerDay := []int{2, 5, 0, 7, 4, 1}
FixBirdCountLog(birdsPerDay)
// => [3 5 1 7 5 1]
*/
func FixBirdCountLog(birdsPerDay []int) []int {
	var slice_index_counter int = 0
	for slice_index_counter = slice_index_counter; slice_index_counter < len(birdsPerDay); slice_index_counter++ {
		if (slice_index_counter+1)%2 != 0 {
			birdsPerDay[slice_index_counter] = (birdsPerDay[slice_index_counter]) + 1
		}
	}
	return birdsPerDay
}

func main() {
	birdsPerDay := []int{3, 0, 5, 1, 0, 4, 1} //nie przechodzi testów więc co trzeci tydzień dodaję by zobaczyć czy tu przejdzie
	fmt.Println(FixBirdCountLog(birdsPerDay))

}

//3 0 5 1 0 4 1 0 3 4 3 0 8 0])
