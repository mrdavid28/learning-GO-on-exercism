package expenses

import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

/*
## 1. Implement a general records filter

Bob deals with a lot of records every day, but not all of them are interesting depending
 on the analysis Bob is making.
Let's help Bob perform some basic filtering of records.

Implement the generic `Filter` function to filter records according to a criteria given by a function.
This filter function accepts a collection of records and a predicate function and returns
 only the records in the collection that satisfy the predicate.

```go
records := []Record{
  {Day: 1, Amount: 15, Category: "groceries"},
  {Day: 11, Amount: 300, Category: "utility-bills"},
  {Day: 12, Amount: 28, Category: "groceries"},
}

// Day1Records only returns true for records that are from day 1
func Day1Records(r Record) bool {
  return r.Day == 1
}

Filter(records, Day1Records)
// =>
// [
//   {Day: 1, Amount: 15, Category: "groceries"}
// ]
```*/
// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var results []Record
	for _, counter := range in {
		if predicate(counter) {
			results = append(results, counter)
		}
	}
	return results
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {

	return func(r Record) bool {
		return p.From <= r.Day && p.To >= r.Day
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(category string) func(Record) bool {
	return func(r Record) bool {
		return category == r.Category
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {

	var sum float64 = 0
	for _, counter := range Filter(in, ByDaysPeriod(p)) {
		sum += counter.Amount
	}
	return sum
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
/***
Calculate the total expenses for records of a category in a period
Use the ByCategory function you made in step 3 to make a category filter function.
Pass that filter function as the second argument of the Filter function you made in step 1.
This allows you to filter the results by a particular category.
After filtering the records by category, check if you have any records for thatcategory.
 you don't, you must return an error.
If you have records belonging to that category,
compute the total expenses for the given period of time using the TotalByPeriod function
 you made in step 4.****/

func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	for _, counter := range Filter(in, ByCategory(c)) {
		if c == counter.Category {
			return TotalByPeriod(Filter(in, ByCategory(c)), p), nil
		}
	}
	return 0, errors.New("nothing")
}
