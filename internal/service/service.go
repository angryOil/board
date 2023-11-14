package service

import (
	"board/internal/domain"
	"board/internal/repository"
	"board/internal/repository/req"
	req2 "board/internal/service/req"
	"board/internal/service/res"
	"context"
	"errors"
	"time"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return Service{repo: repo}
}

const (
	BoardNotFound = "board not found"
)

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

func (s Service) Delete(ctx context.Context, id int) error {
	err := s.repo.Delete(ctx, id)
	return err
}

func (s Service) Patch(ctx context.Context, p req2.Patch) error {
	id := p.Id
	title, content := p.Title, p.Content

	err := s.repo.Patch(ctx, id,
		func(domains []domain.Board) (domain.Board, error) {
			if len(domains) != 1 {
				return domain.NewBuilder().Build(), errors.New(BoardNotFound)
			}
			return domains[0], nil
		},
		func(d domain.Board) (req.Patch, error) {
			u, err := d.Update(title, content)
			if err != nil {
				return req.Patch{}, err
			}
			v := u.ToUpdate()
			return req.Patch{
				Id:        v.Id,
				CafeId:    v.CafeId,
				Writer:    v.Writer,
				BoardType: v.BoardType,
				Title:     v.Title,
				Content:   v.Content,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.LastUpdatedAt,
			}, nil
		},
	)
	return err
}

func (s Service) GetDetail(ctx context.Context, id int) (res.GetDetail, error) {
	domains, err := s.repo.GetDetail(ctx, id)
	if err != nil {
		return res.GetDetail{}, err
	}
	if len(domains) != 1 {
		return res.GetDetail{}, nil
	}
	v := domains[0].ToDetail()
	return res.GetDetail{
		Id:            v.Id,
		BoardType:     v.BoardType,
		Writer:        v.Writer,
		Title:         v.Title,
		Content:       v.Content,
		CreatedAt:     v.CreatedAt,
		LastUpdatedAt: v.LastUpdatedAt,
	}, nil
}
