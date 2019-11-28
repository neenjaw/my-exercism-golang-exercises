// Package clock contains types and functions to represent a clock and manipulate it.
package clock

import "fmt"

// Clock represents the hours and minutes of a clock
type Clock struct {
	hour   int
	minute int
}

// New returns a new clock with the specified hours and minutes normalized to a 24-hr clock.
func New(hh int, mm int) Clock {
	newClock := Clock{hour: 0, minute: 0}
	minutes := hh*60 + mm

	if minutes == 0 {
		return newClock
	}

	if minutes < 0 {
		return newClock.Subtract(minutes * -1)
	}

	return newClock.Add(minutes)
}

// Add uses a positive integer representing minutes to be added to the clock
func (c Clock) Add(mm int) Clock {
	rawMinutes := c.minute + mm
	minute := rawMinutes % 60
	hour := (c.hour + rawMinutes/60) % 24

	return Clock{hour: hour, minute: minute}
}

// Subtract uses a positive integer representing minutes to subtract from the clock.
func (c Clock) Subtract(mm int) Clock {
	rawMinutes := c.minute - mm

	minute := rawMinutes % 60
	hourAdjust := 0
	if minute < 0 {
		minute += 60
		hourAdjust--
	}

	hour := (c.hour + (rawMinutes / 60) + hourAdjust) % 24
	if hour < 0 {
		hour += 24
	}

	return Clock{hour: hour, minute: minute}
}

// String returns a string representation of the Clock struct type, formatted in the hh:mm format
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
