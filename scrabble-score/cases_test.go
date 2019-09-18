package scrabble

// Source: exercism/problem-specifications
// Commit: 0d882ed scrabble-score: Apply new "input" policy
// Problem Specifications Version: 1.1.0

type scrabbleTest struct {
	input    string
	expected int
}

var scrabbleScoreTests = []scrabbleTest{
	{"a", 1},                           // lowercase letter
	{"A", 1},                           // uppercase letter
	{"f", 4},                           // valuable letter
	{"at", 2},                          // short word
	{"zoo", 12},                        // short, valuable word
	{"street", 6},                      // medium word
	{"quirky", 22},                     // medium, valuable word
	{"OxyphenButazone", 41},            // long, mixed-case word
	{"pinata", 8},                      // english-like word
	{"", 0},                            // empty input
	{"abcdefghijklmnopqrstuvwxyz", 87}, // entire alphabet available
}

type scrabbleModTest struct {
	input    string
	mods     []ModifierParam
	expected int
}

var scrabbleScoreModTests = []scrabbleModTest{
	// single letter
	{"a", []ModifierParam{ModifierParam{DoubleLetter, 'a'}}, 2},
	{"A", []ModifierParam{ModifierParam{DoubleLetter, 'A'}}, 2},
	{"a", []ModifierParam{ModifierParam{DoubleLetter, 'A'}}, 2},
	{"A", []ModifierParam{ModifierParam{DoubleLetter, 'a'}}, 2},
	// valuable letter
	{"f", []ModifierParam{ModifierParam{TripleLetter, 'f'}}, 12},
	{"F", []ModifierParam{ModifierParam{TripleLetter, 'F'}}, 12},
	{"f", []ModifierParam{ModifierParam{TripleLetter, 'F'}}, 12},
	{"F", []ModifierParam{ModifierParam{TripleLetter, 'f'}}, 12},
	// short word, multiple mods
	{
		"at",
		[]ModifierParam{
			ModifierParam{DoubleLetter, 'a'},
			ModifierParam{TripleLetter, 't'},
		}, 5},
	// mod for other letter doesn't affect score
	{"F", []ModifierParam{ModifierParam{TripleLetter, 0}}, 4},
	// double word
	{"zoo", []ModifierParam{ModifierParam{DoubleWord, 0}}, 24},
	// triple word
	{"zoo", []ModifierParam{ModifierParam{TripleWord, 0}}, 36},
	// double and triple word
	{"zoo", []ModifierParam{ModifierParam{DoubleWord, 0}, ModifierParam{TripleWord, 0}}, 60},
	// double double word
	{"zoo", []ModifierParam{ModifierParam{DoubleWord, 0}, ModifierParam{DoubleWord, 0}}, 48},
	// triple triple word
	{"zoo", []ModifierParam{ModifierParam{TripleWord, 0}, ModifierParam{TripleWord, 0}}, 72},
	// only one letter bonus
	{"aa", []ModifierParam{ModifierParam{DoubleLetter, 'a'}}, 3},
}
