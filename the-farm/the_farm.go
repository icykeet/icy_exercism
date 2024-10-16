package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fc FodderCalculator, count int) (float64, error) {
	forAll, err1 := fc.FodderAmount(count)
	if err1 != nil {
		return 0, err1
	}
	factor, err2 := fc.FatteningFactor()
	if err2 != nil {
		return 0, err2
	}
	forAllEnd := forAll * factor
	return forAllEnd / float64(count), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, count int) (float64, error) {
	if count <= 0 {
		return 0, errors.New("invalid number of cows")
	} else {
		return DivideFood(fc, count)
	}
}

type InvalidCowsError struct {
	count   int
	message string
}

func (err *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", err.count, err.message)
}

func ValidateNumberOfCows(count int) error {
	var message string
	switch {
	case count < 0:
		message = "there are no negative cows"
	case count == 0:
		message = "no cows don't need food"
	}

	if count <= 0 {
		return &InvalidCowsError{
			count:   count,
			message: message,
		}
	} else {
		return nil
	}
}
