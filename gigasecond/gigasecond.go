// Package gigasecond 
package gigasecond

// import time package from the standard library
import "time"

// Gigasecond represents 1_000_000_000 seconds
const Gigasecond time.Duration = time.Second * 1e9

// AddGigasecond 
func AddGigasecond(t time.Time) time.Time {
	return t.Add(Gigasecond)
}
