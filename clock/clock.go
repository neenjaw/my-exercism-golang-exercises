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
	newClock, minutes := Clock{hour: 0, minute: 0}, hh*60+mm

	(&newClock).makeClock(minutes)

	return newClock
}

func (c *Clock) makeClock(minutes int) {
	rawMinutes := c.minute + minutes

	minute, hourAdjust := rawMinutes%60, 0
	if minute < 0 {
		minute += 60
		hourAdjust--
	}

	hour := (c.hour + (rawMinutes / 60) + hourAdjust) % 24
	if hour < 0 {
		hour += 24
	}

	c.hour, c.minute = hour, minute
}

// Add uses a positive integer representing minutes to be added to the clock
func (c Clock) Add(mm int) Clock {
	newClock := Clock{hour: c.hour, minute: c.minute}

	(&newClock).makeClock(mm)

	return newClock
}

// Subtract uses a positive integer representing minutes to subtract from the clock.
func (c Clock) Subtract(mm int) Clock {
	newClock := Clock{hour: c.hour, minute: c.minute}

	(&newClock).makeClock(mm * -1)

	return newClock
}

// String returns a string representation of the Clock struct type, formatted in the hh:mm format
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
