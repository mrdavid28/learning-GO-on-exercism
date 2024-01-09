package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var sum, loop_counter int = 0, 0

	for loop_counter = 0; loop_counter < len(birdsPerDay); loop_counter++ {
		sum = sum + birdsPerDay[loop_counter]
	}
	return sum
}


// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
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
			/*if slice_size_counter == (slice_size_counter*week)-1 {
				return sum
			}*/

		}
	} else {
		var normal_counter int = (slice_size_counter * week)

		for slice_size_counter = slice_size_counter; slice_size_counter < normal_counter; slice_size_counter++ {
			/*if slice_size_counter == normal_counter {
				return sum
			}*/
			sum = sum + birdsPerDay[slice_size_counter]
		}

	}
	return sum
}
// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	var slice_index_counter int = 0
	for slice_index_counter = slice_index_counter; slice_index_counter < len(birdsPerDay); slice_index_counter++ {
		if (slice_index_counter+1)%2 != 0 {
			birdsPerDay[slice_index_counter] = (birdsPerDay[slice_index_counter]) + 1
		}
	}
	return birdsPerDay
}
