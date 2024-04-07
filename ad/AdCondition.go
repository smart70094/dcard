package ad

type AdCondition struct {
	ID       int
	AgeStart int8
	AgeEnd   int8
	Gender   []string
	Country  []string
	Platform []string
	AdID     int
}
