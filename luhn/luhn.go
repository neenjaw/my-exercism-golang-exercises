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
	luhnSum := 0
	// Use the length to determine whether or not to double the first digit
	doubleDigit := len(input)%2 == 0

	for _, r := range input {
		digit := int(r - '0')

		// encountered an illegal character
		if digit < 0 || digit > 9 {
			return false
		}

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

	return digitCount > 1 && luhnSum%10 == 0
}
