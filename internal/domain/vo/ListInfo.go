package vo

import "time"

type ListInfo struct {
	Id            int
	BoardType     int
	Writer        int
	Title         string
	CreatedAt     time.Time
	LastUpdatedAt time.Time
}
