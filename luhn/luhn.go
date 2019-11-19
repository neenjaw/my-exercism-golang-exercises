// Package luhn has functions for validating a Luhn number.
package luhn

import "strings"

// Valid takes a string representation of a Luhn number and
// performs a validation of it by doubling every second digits,
// starting from the right, and if the doubling is > 9 then 9 is subtracted.
// Finally, if the sum of the digits divides by 10 evenly then it is
// a vaid luhn number.
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	digitCount := 0
	doubleDigit := false
	luhnSum := 0

	// Loop through input in reverse:
	for i := len(input) - 1; i >= 0; i-- {
		r := input[i]

		// encountered an illegal character
		if r < '0' || '9' < r {
			return false
		}

		digit := int(r - '0')

		if doubleDigit {
			digit *= 2
		}
		if digit > 9 {
			digit -= 9
		}

		luhnSum += digit
		digitCount++
		doubleDigit = !doubleDigit
	}

	// If the luhn number is too short
	if digitCount < 2 {
		return false
	}

	// If it's a valid luhn number
	return luhnSum%10 == 0
}
