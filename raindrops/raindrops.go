/*
package raindrops, provides Convert, an exported function to convert an integer to either a string representation of it's factors as rain drops or a string representation of itself.
*/
package raindrops

import (
	"fmt"
	"math"
)

// Convert - a function which takes a number and returns a string
// consisting of up to 3 words.  If the number has 3 as a factor, append
// 'Pling'; if the number has a factor of 5, append 'Plang'; if the number
// has has factor of 7, append Plong.  If it has all three, should
// return them concatenated in that order.  If factors, then return the
// original number as a string.
func Convert(num int) string {
	rain := ""
	rainsounds := []struct {
		factor int
		sound  string
	}{
		{3, "Pling"},
		{5, "Plang"},
		{7, "Plong"},
	}

	for i := range rainsounds {
		if math.Remainder(float64(num), float64(rainsounds[i].factor)) == 0 {
			rain += rainsounds[i].sound
		}
	}

	if rain == "" {
		rain = fmt.Sprintf("%d", num)
	}

	return rain
}
