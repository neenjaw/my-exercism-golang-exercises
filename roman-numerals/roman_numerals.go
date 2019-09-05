/*
Package romannumerals contains functions to convert integers to roman numerals.
*/
package romannumerals

import (
	"errors"
	"strings"
)

var numeralPairs = []struct {
	integer int
	numeral string
}{
	{1000, "M"},
	{900, "CM"},
	{800, "DCCC"},
	{700, "DCC"},
	{600, "DC"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{80, "LXXX"},
	{70, "LXX"},
	{60, "LX"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{8, "VIII"},
	{7, "VII"},
	{6, "VI"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ToRomanNumeral takes an integer between 1-3000 and returns the roman numeral version as a string
func ToRomanNumeral(n int) (string, error) {
	if n < 1 || n > 3000 {
		return "", errors.New("number to be converted must be between 1-3000")
	}

	var numeral strings.Builder

	for _, numeralPair := range numeralPairs {
		for n >= numeralPair.integer {
			numeral.WriteString(numeralPair.numeral)

			n -= numeralPair.integer
		}
	}

	return numeral.String(), nil
}
