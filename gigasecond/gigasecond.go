// Package gigasecond contains functionality for working with the Gigasecond unit of time.
package gigasecond

import (
	"time"
)

// AddGigasecond takes a Time and returns a new Time that is one billion seconds later.
func AddGigasecond(t time.Time) time.Time {
	dur, _ := time.ParseDuration("1000000000s")
	return t.Add(dur)
}
