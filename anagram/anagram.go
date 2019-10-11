package anagram

import (
	"sort"
	"strings"
)

// Detect takes a string subject and a slice of string representing
// possible anagram forms.  Returns a slice of anagrams to the subject
// string.
func Detect(subject string, candidates []string) []string {
	loweredSubject := strings.ToLower(subject)
	sortedSubject := sortString(loweredSubject)
	anagrams := make([]string, 0)

	for _, candidate := range candidates {
		lowered := strings.ToLower(candidate)

		if lowered == loweredSubject {
			continue
		}

		sorted := sortString(lowered)

		if sorted != sortedSubject {
			continue
		}

		anagrams = append(anagrams, candidate)
	}

	return anagrams
}

// Define interface for sorting string by runes
type runeSlice []rune

func (p runeSlice) Len() int           { return len(p) }
func (p runeSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p runeSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func sortString(str string) string {

	runes := []rune(str)

	sort.Sort(runeSlice(runes))

	return string(runes)
}
