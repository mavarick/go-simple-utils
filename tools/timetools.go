package tools

import (
	"time"
)

/* mayor functions:
0, time formatting
1, time offsetting
2, time comparation
3, timestamp to time, vise verse

*/
///////////////////////////////////////////////////////////////////
// time format styles, from  the golang doc
const (
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
)

///////////////////////////////////////////////////////////////////
// time duration format, from golang doc
const (
	Nanosecond  time.Duration = 1
	Microsecond               = 1000 * Nanosecond
	Millisecond               = 1000 * Microsecond
	Second                    = 1000 * Millisecond
	Minute                    = 60 * Second
	Hour                      = 60 * Minute
)

const (
	Day  = 24 * Hour
	Week = 7 * Day
)

type TimeTools struct {
}

func (self *TimeTools) Now() time.Time {
	return time.Now()
}

func (self *TimeTools) Format(t time.Time, format string) string {
	return t.Format(format)
}

func (self *TimeTools) Add(t time.Time, d time.Duration) time.Time {
	return t.Add(d)
}
func (self *TimeTools) Sub(t1, t2 time.Time) time.Duration {
	return t1.Sub(t2)
}
func (self *TimeTools) After(t1, t2 time.Time) bool {
	return t1.After(t2)
}
func (self *TimeTools) Before(t1, t2 time.Time) bool {
	return t1.Before(t2)
}

// format the timestamp to time
func (self *TimeTools) ToTime(sec int64, nsec int64) time.Time {
	return time.Unix(sec, nsec) //notice the time here, not an instance
}

// format the time to timestamp(int64)
func (self *TimeTools) ToTimeStamp(t time.Time) int64 {
	return t.Unix()
}
