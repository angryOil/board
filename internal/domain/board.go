package domain

import (
	"board/internal/domain/vo"
	"errors"
	"time"
)

var _ Board = (*board)(nil)

type Board interface {
	ValidCreate() error
	ValidUpdate() error

	Update(title, content string) (Board, error)

	ToListInfo() vo.ListInfo // 리스트 (간단 정보)
	ToDetail() vo.Detail     // 상세정보 (내용 포함)
}

type board struct {
	id            int
	cafeId        int
	boardType     int
	writer        int
	title         string
	content       string
	createdAt     time.Time
	lastUpdatedAt time.Time
}

const (
	InvalidBoardId   = "invalid board id"
	InvalidCafeId    = "invalid cafe id"
	InvalidBoardType = "invalid board type"
	InvalidWriterId  = "invalid writer id"
	InvalidTitle     = "invalid title"
	InvalidContent   = "invalid content"
	BoardIsNil       = "board is nil"
)

func (b *board) ValidCreate() error {
	if b.cafeId < 1 {
		return errors.New(InvalidCafeId)
	}
	if b.title == "" {
		return errors.New(InvalidTitle)
	}
	if b.content == "" {
		return errors.New(InvalidContent)
	}
	if b.boardType < 1 {
		return errors.New(InvalidBoardType)
	}
	if b.writer < 1 {
		return errors.New(InvalidWriterId)
	}
	return nil
}

func (b *board) ValidUpdate() error {
	if b.id < 1 {
		return errors.New(InvalidContent)
	}
	if b.cafeId < 1 {
		return errors.New(InvalidCafeId)
	}
	if b.title == "" {
		return errors.New(InvalidTitle)
	}
	if b.content == "" {
		return errors.New(InvalidContent)
	}
	if b.boardType < 1 {
		return errors.New(InvalidBoardType)
	}
	if b.writer < 1 {
		return errors.New(InvalidWriterId)
	}
	return nil
}

func (b *board) Update(title, content string) (Board, error) {
	if b == nil {
		return nil, errors.New(BoardIsNil)
	}
	b.title = title
	b.content = content
	err := b.ValidUpdate()
	if err != nil {
		return b, err
	}
	return b, nil
}

func (b *board) ToListInfo() vo.ListInfo {
	return vo.ListInfo{
		Id:            b.id,
		BoardType:     b.boardType,
		Writer:        b.writer,
		Title:         b.title,
		CreatedAt:     b.createdAt,
		LastUpdatedAt: b.lastUpdatedAt,
	}
}

func (b *board) ToDetail() vo.Detail {
	return vo.Detail{
		Id:            b.id,
		BoardType:     b.boardType,
		Writer:        b.writer,
		Title:         b.title,
		Content:       b.content,
		CreatedAt:     b.createdAt,
		LastUpdatedAt: b.lastUpdatedAt,
	}
}
