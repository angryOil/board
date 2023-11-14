package repository

import (
	"board/internal/domain"
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

func (r Repository) Delete(ctx context.Context, id int) error {
	var m model.Board
	_, err := r.db.NewDelete().Model(&m).Where("id = ?", id).Exec(ctx)
	if err != nil {
		log.Println("Delete NewDelete err: ", err)
		return errors.New(InternalServerError)
	}
	return err
}

func (r Repository) Patch(ctx context.Context, id int,
	validFunc func(domains []domain.Board) (domain.Board, error),
	updateFunc func(d domain.Board) (req.Patch, error)) error {

	var models []model.Board
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Println("Patch begin tx err: ", err)
		return errors.New(InternalServerError)
	}
	err = tx.NewSelect().Model(&models).Where("id = ?", id).Scan(ctx)
	if err != nil {
		log.Println("Patch NewSelect err: ", err)
		return errors.New(InternalServerError)
	}

	findDomains := model.ToDomainList(models)

	validDomain, err := validFunc(findDomains)
	if err != nil {
		return err
	}
	updateDto, err := updateFunc(validDomain)
	if err != nil {
		return err
	}
	m := model.ToUpdateModel(updateDto)

	_, err = tx.NewInsert().Model(&m).On("conflict (id) do update").Exec(ctx)
	if err != nil {
		log.Println("Patch NewInsert err: ", err)
		return errors.New(InternalServerError)
	}

	err = tx.Commit()
	if err != nil {
		log.Println("Patch commit err: ", err)
		return errors.New(InternalServerError)
	}
	return nil
}

func (r Repository) GetDetail(ctx context.Context, id int) ([]domain.Board, error) {
	var models []model.Board
	err := r.db.NewSelect().Model(&models).Where("id = ?", id).Scan(ctx)
	if err != nil {
		log.Println("GetDetail NewSelect err: ", err)
		return []domain.Board{}, errors.New(InternalServerError)
	}
	return model.ToDomainList(models), nil
}
