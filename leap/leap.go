/*
Package leap is a library to determine if years are leap years
*/
package leap

// IsLeapYear takes an integer representing a year and returns true if it is a leap year
func IsLeapYear(year int) bool {
	switch {
	case year%400 == 0:
		return true
	case year%100 == 0:
		return false
	case year%4 == 0:
		return true
	default:
		return false
	}
}
