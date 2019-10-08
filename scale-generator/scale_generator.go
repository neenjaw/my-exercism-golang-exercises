// Package scale contains functions that compute the musical scale given a tonic
// and interval.
package scale

import (
	"errors"
	"strings"
)

var sharpTonics = []string{"C", "G", "f#", "a", "A"}
var flatTonics = []string{"F", "bb", "g", "d", "Db", "Eb"}

// Scale fn to return a scale based on a supplied tonic and interval
func Scale(tonic string, interval string) []string {
	if interval == "" {
		interval = "mmmmmmmmmmmm"
	}

	scale := make([]string, 0, 12)

	for note := range generateScale(tonic, interval) {
		scale = append(scale, note)
	}

	return scale
}

// returns a channel which generates a scale based on the tonic and interval
func generateScale(tonic string, interval string) chan string {
	scale, _ := getScale(tonic)

	ch := make(chan string)
	go func() {
		defer close(ch)

		for i, j := 0, 0; j < len(scale); i++ {
			ch <- scale[j]

			if interval[i] == 'm' {
				j++
			} else if interval[i] == 'M' {
				j += 2
			} else if interval[i] == 'A' {
				j += 3
			}
		}
	}()

	return ch
}

// gets the appropriate scale and returns it rotated to the appropriate starting point
// based on the tonic supplied
func getScale(tonic string) ([]string, error) {
	flatTonic := false
	for _, e := range flatTonics {
		if e == tonic {
			flatTonic = true
		}
	}

	var scale []string

	if flatTonic {
		scale = flatScale()
	} else {
		scale = sharpScale()
	}

	rscale, err := rotateScale(scale, tonic)

	if err != nil {
		return nil, err
	}

	return rscale, nil
}

// returns a new slice of strings representing a flat chromatic scale
func flatScale() []string {
	return []string{"F", "Gb", "G", "Ab", "A", "Bb", "B", "C", "Db", "D", "Eb", "E"}
}

// returns a new slice of strings representing a sharp chromatic scale
func sharpScale() []string {
	return []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
}

// takes a slice of strings representing a scale then rotates it to the proper
// starting point based on the tonic
func rotateScale(scale []string, tonic string) ([]string, error) {
	index := -1
	normalizedTonic := strings.Title(tonic)

	for i, note := range scale {
		if note == normalizedTonic {
			index = i
		}
	}

	if index == -1 {
		return nil, errors.New("note must be in scale")
	}

	return append(scale[index:], scale[:index]...), nil
}
