package model

import (
	"board/internal/domain"
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

func ToUpdateModel(u req.Patch) Board {
	return Board{
		Id:            u.Id,
		CafeId:        u.CafeId,
		BoardType:     u.BoardType,
		Writer:        u.Writer,
		Title:         u.Title,
		Content:       u.Content,
		CreatedAt:     u.CreatedAt,
		LastUpdatedAt: u.UpdatedAt,
	}
}

func (m Board) ToDomain() domain.Board {
	return domain.NewBuilder().
		Id(m.Id).
		Writer(m.Writer).
		CafeId(m.CafeId).
		BoardType(m.BoardType).
		Title(m.Title).
		Content(m.Content).
		CreatedAt(m.CreatedAt).
		LastUpdatedAt(m.LastUpdatedAt).
		Build()
}

func ToDomainList(models []Board) []domain.Board {
	result := make([]domain.Board, len(models))
	for i, m := range models {
		result[i] = m.ToDomain()
	}
	return result
}
