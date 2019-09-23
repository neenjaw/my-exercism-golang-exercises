// Package reverse contains functions to reverse strings
package reverse

import (
	"strings"
	"unicode/utf8"
)

// Reverse takes a string and returns the reversed by rune.
func Reverse(s string) string {
	var reverse strings.Builder
	b := []byte(s)

	for utf8.RuneCount(b) > 0 {
		r, size := utf8.DecodeLastRune(b)
		b = b[:len(b)-size]
		reverse.WriteRune(r)
	}

	return reverse.String()
}
