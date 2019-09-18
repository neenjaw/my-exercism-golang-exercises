/*
Package scrabble contains data and functions to calculate the score of a word.

Assumption, the board appears in the traditional layout:

T . . d . . . T . . . d . . T    Where:
. D . . . t . . . t . . . D .      . represents an empty space
. . D . . . d . d . . . D . .      D represents a double word score
d . . D . . . d . . . D . . d      T represents a triple word score
. . . . D . . . . . D . . . .      d represents a double letter score
. t . . . t . . . t . . . t .      t represents a triple letter score
. . d . . . d . d . . . d . .      X represents the center star
T . . d . . . X . . . d . . T
. . d . . . d . d . . . d . .
. t . . . t . . . t . . . t .
. . . . D . . . . . D . . . .
d . . D . . . d . . . D . . d
. . D . . . d . d . . . D . .
. D . . . t . . . t . . . D .
T . . d . . . T . . . d . . T

So it is possible for multiple modifiers to apply to the score at once,
to account for those, the exported Score function has a variadic argument
allowing multiple modifiers to the score to be defined.  The types defined
allow other packages to describe these modifiers.
*/
package scrabble

import "unicode"

// letterValues is a human readable anonymous struct to facilitate a
// the creation of a mapping between the letters and the scores
var letterValues = []struct {
	score   int
	letters string
}{
	{1, "AEIOULNRST"},
	{2, "DG"},
	{3, "BCMP"},
	{4, "FHVWY"},
	{5, "K"},
	{8, "JX"},
	{10, "QZ"},
}

// Do an ETL transform on the letterValues anonymous struct, and return
// a mapping from letter to score.
func makeLetterValues() map[rune]int {
	m := make(map[rune]int)

	for _, letterValue := range letterValues {
		for _, letter := range letterValue.letters {
			m[letter] = letterValue.score
		}
	}

	return m
}

// Modifier is a type to specify if special scoring conditions should be applied
type Modifier int

// Constants which describe the scoring modifier
const (
	DoubleLetter Modifier = iota
	TripleLetter
	DoubleWord
	TripleWord
)

// ModifierParam is a struct to pass scoring modifiers to the Score function
type ModifierParam struct {
	name   Modifier
	letter rune
}

// Score iterates through a string, summing the value of each letter to
// calculate a score.
func Score(word string, mods ...ModifierParam) int {
	letterMods := collectLetterModifiers(mods)
	wordMods := collectWordModifiers(mods)

	letterValues := makeLetterValues()
	initalScoring := 0

	for _, letter := range word {
		upLetter := unicode.ToUpper(letter)

		if c, ok := letterMods[DoubleLetter][upLetter]; ok && c > 0 {
			// double letter score
			initalScoring += (letterValues[upLetter] * 2)
			letterMods[DoubleLetter][upLetter]--

		} else if c, ok := letterMods[TripleLetter][upLetter]; ok && c > 0 {
			// triple letter score
			initalScoring += (letterValues[upLetter] * 3)
			letterMods[TripleLetter][upLetter]--

		} else {
			initalScoring += letterValues[upLetter]
		}
	}

	doubleWordScore := wordMods[DoubleWord] * 2 * initalScoring
	tripleWordScoring := wordMods[TripleWord] * 3 * initalScoring
	finalScoring := doubleWordScore + tripleWordScoring

	if finalScoring == 0 {
		finalScoring = initalScoring
	}

	return finalScoring
}

// collectLetterModifiers processes the variadic parameter from Score and
// looks for letter bonuses, arranged to a map for use in scoring
func collectLetterModifiers(mods []ModifierParam) map[Modifier]map[rune]int {
	letterModifiers := map[Modifier]map[rune]int{
		DoubleLetter: make(map[rune]int),
		TripleLetter: make(map[rune]int),
	}

	for _, mod := range mods {
		mod.letter = unicode.ToUpper(mod.letter)

		if mod.letter != 0 && (mod.name == DoubleLetter || mod.name == TripleLetter) {
			letterModifiers[mod.name][mod.letter]++
		}
	}

	return letterModifiers
}

// collectWordModifiers processes the variadic parameter from Score and
// looks for word bonuses, arranged to a map for use in scoring
func collectWordModifiers(mods []ModifierParam) map[Modifier]int {
	wordModifiers := map[Modifier]int{
		DoubleWord: 0,
		TripleWord: 0,
	}

	for _, mod := range mods {
		if _, ok := wordModifiers[mod.name]; ok {
			wordModifiers[mod.name]++
		}
	}

	return wordModifiers
}
