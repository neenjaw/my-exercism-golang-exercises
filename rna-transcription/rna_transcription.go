/*
Package strand is a package with dna transcription functions
*/
package strand

import "strings"

// ToRNA takes a string representing dna nucleotides, maps them to the rna nucleotide,
// builds the rna string, then returns the completed string.
func ToRNA(dna string) string {
	toRNAMap := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	var rna strings.Builder

	for _, dnaRune := range dna {
		rna.WriteRune(toRNAMap[dnaRune])
	}

	return rna.String()
}
