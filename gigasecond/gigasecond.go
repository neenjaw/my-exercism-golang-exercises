// Package gigasecond performs operations on time
package gigasecond

// import time package from the standard library
import "time"

// Gigasecond represents 1_000_000_000 seconds
const Gigasecond time.Duration = time.Second * 1e9

// AddGigasecond accepts a time as a parameter and adds a gigasecond on to it
func AddGigasecond(t time.Time) time.Time {
	return t.Add(Gigasecond)
}
