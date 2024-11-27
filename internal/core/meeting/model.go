package meeting

import "time"

type Meeting struct {
	ID        string
	Title     string
	StartTime time.Time
	EndTime   time.Time
	Location  string
}
