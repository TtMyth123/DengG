package Bo

import "time"

type BetData struct {
	CycleId string `orm:"size(20)"`
	Num     int
	Money   int
	BetTime time.Time
}
