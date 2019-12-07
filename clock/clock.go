// Package clock contains types and functions to represent a clock and manipulate it.
package clock

import "fmt"

// Clock represents the hours and minutes of a clock
type Clock struct {
	minute int
}

const minutesPerCycle int = 60 * 24

// New returns a new clock with the specified hours and minutes normalized to a 24-hr clock.
func New(hh int, mm int) Clock {
	minute := (hh*60 + mm) % minutesPerCycle
	if minute < 0 {
		minute += minutesPerCycle
	}

	return Clock{minute: minute}
}

// Add uses a positive integer representing minutes to be added to the clock
func (c Clock) Add(mm int) Clock {
	return New(0, c.minute+mm)
}

// Subtract uses a positive integer representing minutes to subtract from the clock.
func (c Clock) Subtract(mm int) Clock {
	return New(0, c.minute-mm)
}

// String returns a string representation of the Clock struct type, formatted in the hh:mm format
func (c Clock) String() string {
	minute, hour := c.minute%60, (c.minute / 60)

	return fmt.Sprintf("%02d:%02d", hour, minute)
}
