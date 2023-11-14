package model

import (
	"board/internal/repository/req"
	"github.com/uptrace/bun"
	"time"
)

type Board struct {
	bun.BaseModel `bun:"table:board,alias:b"`

	Id            int       `bun:"id,pk,autoincrement"`
	CafeId        int       `bun:"cafe_id,notnull"`
	BoardType     int       `bun:"board_type,notnull"`
	Writer        int       `bun:"writer,notnull"`
	Title         string    `bun:"title,notnull"`
	Content       string    `bun:"content,notnull"`
	CreatedAt     time.Time `bun:"created_at"`
	LastUpdatedAt time.Time `bun:"last_updated_at"`
}

func ToCreateModel(c req.Create) Board {
	return Board{
		CafeId:    c.CafeId,
		BoardType: c.BoardType,
		Writer:    c.Writer,
		Title:     c.Title,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
	}
}
