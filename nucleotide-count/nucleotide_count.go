package dna

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram struct {
	counts map[rune]int
}

// DNA is a list of nucleotides.
type DNA []rune

func NewHistogram() Histogram {
	var h Histogram

	h.counts = make(map[rune]int, 4)

	nucleotides := []struct {
		name  rune
		count int
	}{
		{'A', 0},
		{'C', 0},
		{'T', 0},
		{'G', 0},
	}

	for i := range nucleotides {
		h.counts[nucleotides[i].name] = nucleotides[i].count
	}

	return h
}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
//
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	h := NewHistogram()

	return h, nil
}
