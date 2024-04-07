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

type Country string

const (
	Taiwan  Country = "TW"
	Japan   Country = "JP"
	America Country = "US"
)

type Condition struct {
	AgeStart int8       `form:"ageStart" binding:"min=0,max=200"`
	AgeEnd   int8       `form:"AgeEnd" binding:"min=0,max=200"`
	Gender   []Gender   `form:"gender"`
	Country  []Country  `form:"country"`
	Platform []Platform `form:"platform"`
}

type CreateAdReqVo struct {
	Title      string    `form:"title" binding:"required,max=255"`
	StartAt    time.Time `form:"startAt" binding:"required"`
	EndAt      time.Time `form:"endAt" binding:"required"`
	Conditions []Condition
}
