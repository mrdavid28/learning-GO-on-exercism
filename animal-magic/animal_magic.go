package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SeedWithTime seeds math/rand with the current computer time.
//## 1. Seed the random number generator.

//Implement a `SeedWithTime` function that seeds the `math.rand` package with the current computer time.

func SeedWithTime() {
	//panic("Please implement the SeedWithTime function")
	rand.Seed(time.Now().UnixNano())
}

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	rand.Seed(time.Now().UnixNano())
	var d int = (rand.Intn(20) + 1)
	return d
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateEnergyCrypto() float64 {
	/*
	   rand.Seed(time.Now().UnixNano())
	   min := 10
	   max := 30
	   fmt.Println(rand.Intn(max - min + 1) + min)
	*/
	rand.Seed(time.Now().UnixNano())
	//r := min + rand.Float64() * (max - min)
	var f float64 = 0.0 + rand.Float64()*(12.0-0.0)
	return f
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	rand.Seed(time.Now().UnixNano())
	var random_number, loop_counter int = 0, 0
	var randomAnimals = []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	for loop_counter = 0; loop_counter < len(randomAnimals)-1; loop_counter++ {
		random_number = (rand.Intn(7) + 1)
		randomAnimals[loop_counter], randomAnimals[random_number] = randomAnimals[random_number], randomAnimals[loop_counter]
	}
	return randomAnimals
}

func main() {
	//fmt.Println(RollADie())
	//fmt.Println(GenerateEnergyCrypto())

	fmt.Println(ShuffleAnimals())
}
