package main

import (
	"fmt"
)

/*****
## Instructions

In this exercise you'll be working with savings accounts.
Each year, the balance of your savings account is updated based on its interest rate.
The interest rate your bank gives you depends on the amount of money in your account (its balance):

- 3.213% for a balance less than `0` dollars (balance gets more negative).
- 0.5% for a balance greater than or equal to `0` dollars, and less than `1000` dollars.
- 1.621% for a balance greater than or equal to `1000` dollars, and less than `5000` dollars.
- 2.475% for a balance greater than or equal to `5000` dollars.

You have four tasks, each of which will deal your balance and its interest rate.
******/

/*****
## 1. Calculate the interest rate

Implement the `InterestRate()` function to calculate the interest rate based on the specified balance:

```go
InterestRate(200.75)
// => 0.5
```*****/
// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var interest_rate float32 = 0
	if balance < 0 {
		interest_rate = 3.213
	} else if balance >= 0 && balance < 1000 {
		interest_rate = 0.5
	} else if balance >= 1000 && balance < 5000 {
		interest_rate = 1.621
	} else {
		interest_rate = 2.475
	}
	return interest_rate
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	var Interest float64 = 0
	if balance < 0 {
		Interest = 0.03213 * balance
	} else if balance >= 0 && balance < 1000 {
		Interest = 0.005 * balance
	} else if balance >= 1000 && balance < 5000 {
		Interest = 0.01621 * balance
	} else {
		Interest = 0.02475 * balance
	}
	return Interest

}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
/**## 3. Calculate the annual balance update

Implement the `AnnualBalanceUpdate()` function to calculate the annual balance update, taking into account the interest rate:

```go
AnnualBalanceUpdate(200.75)
// => 201.75375
```**/
func AnnualBalanceUpdate(balance float64) float64 {
	var Interest, AnnualBalanceUpdate float64 = 0, 0
	if balance < 0 {
		Interest = 0.03213 * balance
		AnnualBalanceUpdate = balance + Interest
	} else if balance >= 0 && balance < 1000 {
		Interest = 0.005 * balance
		AnnualBalanceUpdate = balance + Interest
	} else if balance >= 1000 && balance < 5000 {
		Interest = 0.01621 * balance
		AnnualBalanceUpdate = balance + Interest
	} else {
		Interest = 0.02475 * balance
		AnnualBalanceUpdate = balance + Interest
	}
	return AnnualBalanceUpdate
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
/*Implement the `YearsBeforeDesiredBalance()` function to calculate the minimum number of years required to reach the desired balance, taking into account
that each year, interest is added to the balance.
This means that the balance after one year is: start balance + interest for start balance.
The balance after the second year is: balance after one year + interest for balance after one year.
And so on, until the current year's balance is greater than or equal to the target balance.*/

func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var Interest float64 = 0
	var years int = 0
	for years = 0; balance < targetBalance; years++ {
		if balance < 0 {
			Interest = 0.03213 * balance
			balance = balance + Interest
		} else if balance >= 0 && balance < 1000 {
			Interest = 0.005 * balance
			balance = balance + Interest
		} else if balance > 1000 && balance < 5000 {
			Interest = 0.01621 * balance
			balance = balance + Interest
		} else {
			Interest = 0.02475 * balance
			balance = balance + Interest
		}
	}
	return years
}

func main() {
	fmt.Println(YearsBeforeDesiredBalance(1000.000000, 1032.682765))
}
