package isogram

import "unicode"

// IsIsogram tests a string and if there are no repeating letters returns true.
func IsIsogram(s string) bool {
	lettersUsed := make(map[rune]bool, 26)

	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}

		lowered := unicode.ToLower(r)

		if lettersUsed[lowered] {
			return false
		}

		lettersUsed[lowered] = true
	}

	return true
}
