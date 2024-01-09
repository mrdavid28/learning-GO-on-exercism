package thefarm

import (
	"errors"
	"strconv"
)

type InvalidCowsError struct {
	number_of_cows int
	message        string
}

func (err *InvalidCowsError) Error() string {
	var error_message string
	error_message = strconv.Itoa(err.number_of_cows) + " cows are invalid:" + err.message
	
	return error_message
}

func DivideFood(fodder_calc FodderCalculator, number_of_cows int) (float64, error) {

	fodder_amount, err := fodder_calc.FodderAmount(number_of_cows)
	if err != nil {
		return 0, err
	}
	fattening_factor, err := fodder_calc.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return fodder_amount * fattening_factor / float64(number_of_cows), nil

} // TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodder_calc FodderCalculator, number_of_cows int) (float64, error) {

	if number_of_cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fodder_calc, number_of_cows)

}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(number_of_cows int) error {
	if number_of_cows == 0 {
		return &InvalidCowsError{
			message:        " no cows don't need food",
			number_of_cows: 0,
		}
	} else if number_of_cows < 0 {
		return &InvalidCowsError{
			message:        " there are no negative cows",
			number_of_cows: number_of_cows,
		}
	}

	return nil

}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain panic("").
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
