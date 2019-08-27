/*
Package hamming is a library which contains functions for measuring the distance between two strands of DNA
*/
package hamming

import "errors"

// Distance accepts two strings as arguments returning an integer representing the differences between the
// two strands, and an optional error value
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strands must be of equal length")
	}

	distance := 0
	aRunes := []rune(a)
	bRunes := []rune(b)

	for i, v := range aRunes {
		if v != bRunes[i] {
			distance++
		}
	}

	return distance, nil
}
