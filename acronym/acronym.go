// Package acronym has functions to create acronyms
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate takes a string as input, returns an upcased abbreviation.
func Abbreviate(s string) string {
	var abbr strings.Builder

	take := true
	for _, r := range s {
		switch {
		case take && unicode.IsLetter(r):
			take = false
			abbr.WriteRune(unicode.ToUpper(r))
		case r == ' ' || r == '-':
			take = true
		}
	}

	return abbr.String()
}
