package pangram

import "unicode"

const alphabetLength = 26

// IsPangram tests a string and if it uses all of the letter of the
// alphabet, function returns true.
func IsPangram(s string) bool {
	lettersUsed := make(map[rune]bool, alphabetLength)

	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}

		lowered := unicode.ToLower(r)
		lettersUsed[lowered] = true
	}

	return len(lettersUsed) == alphabetLength
}
