/*
Package space is a library to convert an age in earth seconds to a specified planet's years.
*/
package space

// Planet type is a string representing a planet
type Planet string

var secondsPerEarthYear = 31557600.0

var yearRelativeToEarthYear = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age is takes an age in earth seconds and the name of a Planet,
// which it then converts to the specified planet's age in years
func Age(seconds float64, planet Planet) float64 {
	return seconds / secondsPerEarthYear / yearRelativeToEarthYear[planet]
}
