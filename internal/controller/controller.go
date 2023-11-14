package controller

import (
	"board/internal/controller/req"
	"board/internal/service"
	req2 "board/internal/service/req"
	"context"
)

type Controller struct {
	s service.Service
}

func NewController(s service.Service) Controller {
	return Controller{s: s}
}

func (c Controller) Create(ctx context.Context, cafeId int, boardType int, dto req.Create) error {
	err := c.s.Create(ctx, req2.Create{
		CafeId:    cafeId,
		BoardType: boardType,
		Writer:    dto.Writer,
		Title:     dto.Title,
		Content:   dto.Content,
	})
	return err
}

func (c Controller) Delete(ctx context.Context, id int) error {
	err := c.s.Delete(ctx, id)
	return err
}

func (c Controller) Patch(ctx context.Context, id int, p req.Patch) error {
	err := c.s.Patch(ctx, req2.Patch{
		Id:      id,
		Title:   p.Title,
		Content: p.Content,
	})
	return err
}
