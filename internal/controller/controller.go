package controller

import (
	"board/internal/controller/req"
	"board/internal/controller/res"
	page2 "board/internal/page"
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

func (c Controller) GetDetail(ctx context.Context, id int) (res.Detail, error) {
	d, err := c.s.GetDetail(ctx, id)
	if err != nil {
		return res.Detail{}, err
	}
	return res.Detail{
		Id:            d.Id,
		BoardType:     d.BoardType,
		Writer:        d.Writer,
		Title:         d.Title,
		Content:       d.Content,
		CreatedAt:     d.CreatedAt,
		LastUpdatedAt: d.LastUpdatedAt,
	}, nil
}

func (c Controller) GetList(ctx context.Context, cafeId int, boardType int, writer int, reqPage page2.ReqPage) ([]res.Info, int, error) {
	list, total, err := c.s.GetListTotal(ctx, cafeId, boardType, writer, reqPage)
	if err != nil {
		return []res.Info{}, 0, err
	}
	dto := make([]res.Info, len(list))
	for i, l := range list {
		dto[i] = res.Info{
			Id:            l.Id,
			BoardType:     l.BoardType,
			Writer:        l.Writer,
			Title:         l.Title,
			CreatedAt:     l.CreatedAt,
			LastUpdatedAt: l.LastUpdatedAt,
		}
	}
	return dto, total, err
}
