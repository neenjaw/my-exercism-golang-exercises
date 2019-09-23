// Package darts contains functions to calculate dart scoring
package darts

import "math"

var circles = []struct {
	r      float64
	points int
}{
	{1.0, 10},
	{5.0, 5},
	{10.0, 1},
}

// Score takes an x, y coordinate pair then computes a score
func Score(x, y float64) int {
	points := 0

	for _, c := range circles {
		if inCircle(c.r, x, y) {
			points = c.points
			break
		}
	}

	return points
}

// Given the formula to find the distance between two points:
//   d = sqrt((Xa-Xb)^2 + (Ya-Yb)^2)
// we can use the center point of the circle as the second point:
//   d = sqrt((x - 0)^2 + (y - 0)^2)
// and simplify to determine if a point is inside a circle of radius r:
//   x^2 + y^2 <= r^2
func inCircle(r, x, y float64) bool {
	return math.Pow(x, 2.0)+math.Pow(y, 2.0) <= math.Pow(r, 2.0)
}
