package dna

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram struct {
	counts map[rune]int
}

// DNA is a list of nucleotides.
type DNA []rune

// Init is a method of type Histogram to initialize the nucleotide counts
func (h Histogram) Init() {
	h.counts = map[rune]int{
		'A': 0,
		'C': 0,
		'T': 0,
		'G': 0,
	}
}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
//
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	h := new(Histogram)
	h.Init()

	return *h, nil
}
