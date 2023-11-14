package vo

import "time"

type Detail struct {
	Id            int
	BoardType     int
	Writer        int
	Title         string
	Content       string
	CreatedAt     time.Time
	LastUpdatedAt time.Time
}
