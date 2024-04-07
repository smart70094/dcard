package ad

import "time"

type HotAd struct {
	Title   string
	StartAt time.Time
	EndAt   time.Time
}
