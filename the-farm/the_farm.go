package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	amt int
	msg string
}

func (err InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", err.amt, err.msg)
}

// TODO: define the 'DivideFood' function
func DivideFood(calc FodderCalculator, cows int) (float64, error) {
	amt, err := calc.FodderAmount(cows)
	if err != nil {
		return 0.0, err
	}
	factor, err := calc.FatteningFactor()
	if err != nil {
		return 0.0, err
	}
	return amt * factor / float64(cows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(calc FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0.0, errors.New("invalid number of cows")
	}
	return DivideFood(calc, cows)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{cows, "there are no negative cows"}
	} else if cows == 0 {
		return &InvalidCowsError{cows, "no cows don't need food"}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
