package encode

import (
	"strconv"
	"strings"
	"unicode"
)

// RunLengthEncode takes a message and replaces repeated characters with a string integer
func RunLengthEncode(message string) string {
	var encoded strings.Builder

	first := true
	var lastRune rune
	var count int
	for _, r := range message {
		// On first iteration, set the initial state
		if first {
			lastRune = r
			count = 1
			first = false
			continue
		}

		// If the runes match, then increase count and move on
		if r == lastRune {
			count++
			continue
		}

		// Runes don't match, so write previous to builder
		if count > 1 {
			encoded.WriteString(strconv.Itoa(count))
		}
		encoded.WriteRune(lastRune)

		// reset for next
		lastRune = r
		count = 1
	}

	// Do last write
	if count > 1 {
		encoded.WriteString(strconv.Itoa(count))
	}
	// check is done so that null byte isn't written to the builder
	if !first {
		encoded.WriteRune(lastRune)
	}

	return encoded.String()
}

// RunLengthDecode takes an encoded string, where repeated characters are represented by numbers and returns the full decoded string
func RunLengthDecode(message string) string {
	var decoded strings.Builder

	numbersAsRunes := []rune{}
	for _, r := range message {
		// collect the numbers
		if unicode.IsNumber(r) {
			numbersAsRunes = append(numbersAsRunes, r)
			continue
		}

		// if no numbers collected, just write the single rune
		if len(numbersAsRunes) == 0 {
			decoded.WriteRune(r)
			continue
		}

		// take the collected numbers, then repeat the run n times
		count, _ := strconv.Atoi(string(numbersAsRunes))
		for i := 0; i < count; i++ {
			decoded.WriteRune(r)
		}

		// reset for next
		numbersAsRunes = []rune{}
	}

	return decoded.String()
}
