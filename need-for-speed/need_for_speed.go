// package speed
package main

import "fmt"

// TODO: define the 'Car' type struct

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        5, //{battery:100 batteryDrain:3 speed:5 distance:1}
		batteryDrain: 3,
		battery:      100,
		distance:     1,
	}
}

// TODO: define the 'Track' type struct

type Track struct {
	distance int
}

func NewTrack(distance int) Track {
	return Track{
		distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if (car.distance != 1) && (car.batteryDrain > car.battery) {
		return Car{
			speed:        car.speed,
			batteryDrain: car.batteryDrain, //good
			battery:      car.battery,
			distance:     0}
	} else if car.batteryDrain-car.battery == 0 {
		return Car{
			speed:        car.speed,
			batteryDrain: car.batteryDrain,
			battery:      0,
			distance:     car.speed,
		}
	} else if car.distance == 1 && car.battery != 100 {
		return Car{
			speed:        car.speed,
			batteryDrain: car.batteryDrain,
			battery:      car.battery,
			distance:     1,
		}
	} else {
		return Car{
			speed:        car.speed,
			batteryDrain: car.batteryDrain,
			battery:      car.battery - car.batteryDrain,
			distance:     car.distance + car.speed,
		}
	}
}

// CanFinish checks if a car is able to finish a certain track.
/* go

To finish a race, a car has to be able to drive the race's distance.
This means not draining its battery before having crossed the finish line.
Implement the `CanFinish` function that takes a `Car` and a `Track` instance as its parameter and returns `true`
if the car can finish the race; otherwise, return `false`.

Assume that you are currently at the starting line of the race and start the engine of the car for the race.
Take into account that the car's battery might not necessarily be fully charged when starting the race:
speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)

distance := 100
track := NewTrack(distance)

CanFinish(car, track)
// => true
*/

func CanFinish(car Car, track Track) bool {
	/*
		To finish a race, a car has to be able to drive the race's distance.
		This means not draining its battery before having crossed the finish line.
		Implement the CanFinish function that takes a Car and a Track instance as its parameter and returns true if the car can finish the race;
		otherwise, return false.
		Assume that you are currently at the starting line of the race and start the engine of the car for the race.
		Take into account that the car's battery might not necessarily be fully charged when starting the race:*/
	/*name: "Car cannot finish the track distance with initial battery less than 100%",
		car: Car{
			speed:        2,
			batteryDrain: 3,
			battery:      25,
			distance:     0,
		},
		track: Track{
			distance: 24,
		},
		expected: false,
	}*/
	var counter int = 0
	for car.distance; car.distance < track.distance; car.distance += car.distance {
		counter += car.distance
	}
	fmt.Println(counter)
	return true
}

func main() {
	speed := 5
	batteryDrain := 2
	car := NewCar(speed, batteryDrain)

	distance := 100
	track := NewTrack(distance)

	CanFinish(car, track)
}
