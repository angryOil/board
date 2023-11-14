package req

import "time"

type Patch struct {
	Id        int
	CafeId    int
	Writer    int
	BoardType int
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
