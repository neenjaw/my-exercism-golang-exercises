/*
Package triangle - a library to determine the kind of triangle based on input of the sides
*/
package triangle

import "sort"
import "math"

// Kind - referring to the kind of triangle
type Kind int

// NaT - Not a triangle
// Equ - An Equilateral triangle
// Iso - An Isosceles triangle
// Sca - A Scalene triangle
const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

// KindFromSides takes 3 numbers and returns the Kind of a triangle is formed by the three sides
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	sides := []float64{a, b, c}
	sort.Float64s(sides)

	switch {
	case hasInvalidSide(sides), !isClosed(sides):
		k = NaT
	case isEquilateral(sides):
		k = Equ
	case isIsosceles(sides):
		k = Iso
	default:
		k = Sca
	}

	return k
}

func hasInvalidSide(sortedSides []float64) bool {
	hasSideLessZero := sortedSides[0] <= 0
	hasInfiniteSide := math.IsInf(sortedSides[2], 1)

	return hasSideLessZero || hasInfiniteSide
}

func isClosed(sortedSides []float64) bool {
	return sortedSides[0]+sortedSides[1] >= sortedSides[2]
}

func isEquilateral(sortedSides []float64) bool {
	for i := 1; i < len(sortedSides); i++ {
		if sortedSides[i] != sortedSides[0] {
			return false
		}
	}

	return true
}

func isIsosceles(sortedSides []float64) bool {
	shorterSidesEqual := sortedSides[0] == sortedSides[1]
	longerSidesEqual := sortedSides[1] == sortedSides[2]

	return shorterSidesEqual || longerSidesEqual
}
