// package partyrobot
package main

import (
	"fmt"
	"strconv"
)

// Welcome greets a person by name.
/*
Implement the `Welcome` function to return a welcome message using the given name:

```go
Welcome("Christiane")
// => Welcome to my party, Christiane!
```*/
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
/*Welcome a new guest to the party whose birthday is today

Implement the `HappyBirthday` function to return a birthday message using the given name and age of the person.
Unfortunately the programmer is a bit of a show-off, so the robot has to demonstrate its knowledge of every guest's birthday.

```go
HappyBirthday("Frank", 58)
// => Happy birthday Frank! You are now 58 years old!*/
func HappyBirthday(name string, age int) string {
	return fmt.Sprint("Happy birthday " + name + "! You are now " + strconv.Itoa(age) + " years old!")
}

/*
## 3. Give directions

Implement the `AssignTable` function to give directions.
It should accept 5 parameters.

- The name of the new guest to greet (`string`)
- The table number (`int`)
- The name of the seatmate (`string`)
- The direction where to find the table (`string`)
- The distance to the table (`float64`)

The exact result format can be seen in the example below.

The robot provides the table number in a 3 digits format.
If the number is less than 3 digits it gets extra leading zeroes to become 3 digits (e.g. 3 becomes 003).
The robot also mentions the distance of the table.
Fortunately only with a precision that's limited to 1 digit.

```go
AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298)
// =>
// Welcome to my party, Christiane!
// You have been assigned to table 027. Your table is left, exactly 23.8 meters from here.
// You will be sitting next to Frank.*/

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", name, table, direction, distance, neighbor)
}

func main() {
	fmt.Println(AssignTable("Dawid", 92, "Paula", "center", 77.7))
}
