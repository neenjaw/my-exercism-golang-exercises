// Package luhn has functions for validating a Luhn number.
package luhn

import (
	"errors"
)

// Valid takes a string representation of a Luhn number and
// performs a validation of it by doubling every second digits,
// starting from the right, and if the doubling is > 9 then 9 is subtracted.
// Finally, if the sum of the digits divides by 10 evenly then it is
// a vaid luhn number.
func Valid(input string) bool {
	numbers, err := getDigits(input)

	// Something invalid with the input, therefore not luhn
	if err != nil {
		return false
	}

	// Single digits can't be luhn numbers
	if len(numbers) == 1 {
		return false
	}

	// loop through the digits in reverse starting with the second last one
	for i := len(numbers) - 2; i >= 0; i = i - 2 {
		double := numbers[i] * 2

		if double > 9 {
			double = double - 9
		}

		numbers[i] = double
	}

	sum := 0
	for _, number := range numbers {
		sum += number
	}

	if sum%10 == 0 {
		return true
	}

	return false
}

func getDigits(input string) ([]int, error) {
	numbers := make([]int, 0)

	for _, r := range input {
		// if the rune is a digit, then append the digit's value to the slice
		if '0' <= r && r <= '9' {
			numbers = append(numbers, int(r-'0'))
			continue
		}

		// Ignore whitespace
		if r == ' ' {
			continue
		}

		// if there is any other rune, return an error.
		return nil, errors.New("invalid luhn input")
	}

	return numbers, nil
}
