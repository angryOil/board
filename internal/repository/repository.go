package repository

import (
	"board/internal/repository/model"
	"board/internal/repository/req"
	"context"
	"errors"
	"github.com/uptrace/bun"
	"log"
)

type Repository struct {
	db bun.IDB
}

func NewRepository(db bun.IDB) Repository {
	return Repository{db: db}
}

const (
	InternalServerError = "internal server error"
)

func (r Repository) Create(ctx context.Context, c req.Create) error {
	m := model.ToCreateModel(c)
	_, err := r.db.NewInsert().Model(&m).Exec(ctx)

	if err != nil {
		log.Println("Create NewInsert err: ", err)
		return errors.New(InternalServerError)
	}
	return nil
}
