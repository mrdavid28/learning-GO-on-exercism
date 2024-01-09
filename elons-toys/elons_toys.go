package elon

import "strconv"

// TODO: define the 'Drive()' method
/*## 1. Drive the car

Implement the `Drive` method on the `Car` that updates the number of meters driven based on the car's speed, and reduces the battery according to the battery drainage:

```go
speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)
car.Drive()
// car is now Car{speed: 5, batteryDrain: 2, battery: 98, distance: 5}
```

Note: You should not try to drive the car if doing so will cause the car's battery to be below 0. */
func (car *Car) Drive() {
	if car.batteryDrain > car.battery {
		car.distance += 0
		car.battery += 0
	} else {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}

}

// TODO: define the 'DisplayDistance() string' method
/*## 2. Display the distance driven

Implement a `DisplayDistance` method on `Car` to return the distance
as displayed on the LED display as a `string`:

```go
speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)

fmt.Println(car.DisplayDistance())
// Output: "Driven 0 meters"
```*/
func (car *Car) DisplayDistance() string {
	var driven_distance string = "Driven " + strconv.Itoa(car.distance) + " meters"
	return driven_distance
}

// TODO: define the 'DisplayBattery() string' method
/*## 3. Display the battery percentage

Implement the `DisplayBattery` method on `Car` to return the battery percentage as displayed on the LED display as a `string`:

```go
speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)

fmt.Println(car.DisplayBattery())
// Output: "Battery at 100%"
```*/

func (car *Car) DisplayBattery() string {
	var battery_level string = "Battery at " + strconv.Itoa(car.battery) + "%"
	return battery_level
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
/*## 4. Check if a remote control car can finish a race
To finish a race, a car has to be able to drive the race's distance.
 This means not draining its battery before having crossed the finish line.
 Implement the `CanFinish` method that takes a `trackDistance int`
 as its parameter and returns `true` if the car can finish the race; otherwise, return `false`:

```go
speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)

trackDistance := 100

car.CanFinish(trackDistance)
// => true
```*/
func (car *Car) CanFinish(trackDistance int) bool {
	for car.battery >= car.batteryDrain {
		car.Drive()
	}
	return car.distance >= trackDistance
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
