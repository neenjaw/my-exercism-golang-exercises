// Package acronym
package acronym

import "strings"
import "unicode"

// Abbreviate should have a comment documenting it.
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
