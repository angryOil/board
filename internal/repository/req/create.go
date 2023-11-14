package req

import "time"

type Create struct {
	CafeId    int
	Writer    int
	BoardType int
	Title     string
	Content   string
	CreatedAt time.Time
}
