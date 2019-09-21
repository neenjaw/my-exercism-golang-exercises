/*
Package protein contains functions to translate strings
representing RNA into amino acids as part of protein
transcription.
*/
package protein

import "fmt"

// ErrStop - reached end protein signal
var ErrStop = fmt.Errorf("stop message encountered")

// ErrInvalidBase - Malformed codon
var ErrInvalidBase = fmt.Errorf("malformed or bad codon encountered")

// intermediary type to match codons to amino acids
type message int

const (
	methionine message = iota
	phenylalanine
	leucine
	serine
	tyrosine
	tryptophan
	cysteine
	stop
)

var codonToMessage = map[string]message{
	"AUG": methionine,
	"UUU": phenylalanine,
	"UUC": phenylalanine,
	"UUA": leucine,
	"UUG": leucine,
	"UCU": serine,
	"UCC": serine,
	"UCA": serine,
	"UCG": serine,
	"UAU": tyrosine,
	"UAC": tyrosine,
	"UGU": cysteine,
	"UGC": cysteine,
	"UGG": tryptophan,
	"UAA": stop,
	"UAG": stop,
	"UGA": stop,
}

var messageToString = map[message]string{
	methionine:    "Methionine",
	phenylalanine: "Phenylalanine",
	leucine:       "Leucine",
	serine:        "Serine",
	tyrosine:      "Tyrosine",
	tryptophan:    "Tryptophan",
	cysteine:      "Cysteine",
}

// FromCodon takes a string representing a codon and returns the aminoacid
func FromCodon(codon string) (string, error) {
	if m, ok := codonToMessage[codon]; ok {
		if m == stop {
			return "", ErrStop
		}

		return messageToString[m], nil
	}

	return "", ErrInvalidBase
}

// FromRNA takes a string of RNA, returns the amino acids
func FromRNA(rna string) ([]string, error) {
	aminos := make([]string, 0)

translationLoop:
	for n := 0; n < len(rna); n += 3 {
		if n+3 > len(rna) {
			return nil, ErrInvalidBase
		}

		m, err := FromCodon(rna[n : n+3])

		switch err {
		case ErrStop:
			break translationLoop
		case ErrInvalidBase:
			return aminos, ErrInvalidBase
		default:
			aminos = append(aminos, m)
		}
	}

	return aminos, nil
}
