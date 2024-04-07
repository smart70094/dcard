package ad

import "time"

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
)

type Platform string

const (
	Android Platform = "android"
	IOS     Platform = "ios"
	Web     Platform = "web"
)

var CountryMap = map[string]bool{
	"TW": true,
	"JP": true,
	"US": true,
}

type Condition struct {
	AgeStart int8       `form:"ageStart" binding:"min=0,max=200"`
	AgeEnd   int8       `form:"AgeEnd" binding:"min=0,max=200"`
	Gender   []Gender   `form:"gender"`
	Country  []string   `form:"country"`
	Platform []Platform `form:"platform"`
}

type CreateAdReqVo struct {
	Title      string    `form:"title" binding:"required,max=255"`
	StartAt    time.Time `form:"startAt" binding:"required"`
	EndAt      time.Time `form:"endAt" binding:"required"`
	Conditions []Condition
}
