package main

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
/*## 1. Parse appointment date

Implement the `Schedule` function to parse a textual representation of an appointment date into the corresponding `time.Time` format:

```going
Schedule("7/25/2019 13:45:00")
// => 2019-07-25 13:45:00 +0000 UTC


func parseTime() time.Time {
    date := "Tue, 09/22/1995, 13:00"
    layout := "Mon, 01/02/2006, 15:04"

    t := time.Parse(layout,date) // time.Time, error
}

// => 1995-09-22 13:00:00 +0000 UTC*/
func Schedule(date string) time.Time {
	var layout string = "1/2/2006 15:04:05" // there is year, month and day in that format also it must reference date 1/2/2006 15:04:05
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed
/*## 2. Check if an appointment has already passed

Implement the `HasPassed` function that takes an appointment date and checks if the appointment was somewhere in the past:

```go
HasPassed("July 25, 2019 13:45:00")
// => true*/
func HasPassed(date string) bool {
	//September 15th in 2012.
	appointment, _ := time.Parse("1/2/2006 15:04:05", date)
	var actual_date time.Time = time.Now()
	return appointment.Before(actual_date)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	isafternoon, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	var hour int = isafternoon.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	appointment, _ := time.Parse("1/2/2006 15:04:05", date)
	return appointment.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
// In this exercise you'll be working on an appointment scheduler for a beauty salon that opened on SEPTEMBER 15TH in 2012.
func AnniversaryDate() time.Time {
	var t time.Time = time.Now().Year()
	var year time.Time = t.Year()
	return year
}

func main() {
	//.Println("pa")
	var date string = "Friday, September 9, 2112 11:59:59"
	fmt.Println(Description(date))
}
