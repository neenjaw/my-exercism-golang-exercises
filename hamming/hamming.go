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

	return getDistance(a, b, 0), nil
}

// recursively compare the first character of the string, return once count complete
func getDistance(a, b string, count int) int {
	if a == "" {
		return count
	}

	if a[:1] == b[:1] {
		return getDistance(a[1:], b[1:], count)
	}

	return getDistance(a[1:], b[1:], count+1)
}
