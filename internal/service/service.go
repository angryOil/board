package service

import (
	"board/internal/domain"
	"board/internal/repository"
	"board/internal/repository/req"
	req2 "board/internal/service/req"
	"context"
	"time"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return Service{repo: repo}
}

func (s Service) Create(ctx context.Context, c req2.Create) error {
	cafeId, writer, boardType := c.CafeId, c.Writer, c.BoardType
	title, content := c.Title, c.Content
	createdAt := time.Now()

	err := domain.NewBuilder().
		CafeId(cafeId).
		Writer(writer).
		BoardType(boardType).
		Title(title).
		Content(content).
		CreatedAt(createdAt).
		Build().ValidCreate()
	if err != nil {
		return err
	}

	err = s.repo.Create(ctx, req.Create{
		CafeId:    cafeId,
		Writer:    writer,
		BoardType: boardType,
		Title:     title,
		Content:   content,
		CreatedAt: createdAt,
	})
	return err
}
