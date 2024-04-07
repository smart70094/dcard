package ad

import "time"

type Ad struct {
	Title   string
	StartAt time.Time
	EndAt   time.Time
}
