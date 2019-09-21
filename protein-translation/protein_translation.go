package protein

// ErrStop indicates a codon error.
type ErrStop struct {
	s string
}

func (e *ErrStop) Error() string {
	return e.s
}

// ErrInvalidBase indicates an invalid codon pattern
type ErrInvalidBase struct {
	s string
}

func (e *ErrInvalidBase) Error() string {
	return e.s
}

// intermediary type
type message int

const (
	methionine message = iota
	phenylalanine
	leucine
	serine
	tyrosine
	tryptophan
	stop
)

var codonMessage = map[string]message{
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
}

func FromCodon(codon string) (string, error) {
	if m, ok := codonMessage[codon]; ok {
		if m == stop {
			return "", &ErrStop{"stop encountered"}
		}

		return messageToString[m], nil
	}

	return "", &ErrInvalidBase{"invalid sequence"}
}

func FromRNA(rna string) ([]string, error) {
	aminos := make([]string)
	for n := 0; n < len(rna); n += 3 {
		if n+2 > len(rna) {
			return "", &ErrInvalidBase{"incomplete trailing a codon"}
		}

		if m, err := FromCodon(rna[n : n+2]); err != nil {
			if _, ok := err.(*ErrStop); ok {
				break
			}

			if _, ok := err.(*ErrInvalidBase); ok {
				return "", &ErrInvalidBase{"invalid codon"}
			}

			aminos = append(aminos, messageToString[m])
		}
	}

	return aminos, nil
}
