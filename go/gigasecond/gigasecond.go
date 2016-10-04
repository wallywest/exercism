// Package clause.
package gigasecond

import "time"

// Constant declaration.
const testVersion = 4 // find the value in gigasecond_test.go

const gigasecond = 1000000000 * time.Second

// API function.  It uses a type from the Go standard library.
func AddGigasecond(in time.Time) (out time.Time) {
	return in.Add(gigasecond)
}
