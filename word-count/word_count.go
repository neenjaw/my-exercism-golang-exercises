// Package wordcount is a package containing a frequency type and function
// to examine a phrase.
package wordcount

import (
	"regexp"
	"strings"
)

// Frequency represents a map where the string key returns an integer value
type Frequency map[string]int

// WordCount takes a phrase, finds all the words excluding whitespace and
// punctuation, returns a Frequency containing each word and it's associated
// count.
func WordCount(phrase string) Frequency {
	wordCounts := make(Frequency, 40)

	splitIntoWords := regexp.MustCompile("[[:alnum:]]+('t)?")
	words := splitIntoWords.FindAllString(phrase, -1)

	for _, word := range words {
		if word == "" {
			continue
		}

		wordCounts[strings.ToLower(word)]++
	}

	return wordCounts
}
