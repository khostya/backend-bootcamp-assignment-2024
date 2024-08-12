package repo

import (
	"backend-bootcamp-assignment-2024/internal/domain"
	"backend-bootcamp-assignment-2024/internal/repo/exec"
	"backend-bootcamp-assignment-2024/internal/repo/schema"
	"backend-bootcamp-assignment-2024/internal/repo/transactor"
	"context"
	sq "github.com/Masterminds/squirrel"
	"time"
)

const (
	houseTable = "bootcamp.houses"
)

type (
	House struct {
		queryEngineProvider transactor.QueryEngineProvider
	}
)

func NewHouseRepo(provider transactor.QueryEngineProvider) House {
	return House{provider}
}

func (h House) Create(ctx context.Context, house domain.House) (uint, error) {
	db := h.queryEngineProvider.GetQueryEngine(ctx)

	record := schema.NewHouse(house)

	query := sq.Insert(houseTable).
		Columns(record.ColumnsInsert()...).
		Values(record.ValuesInsert()...).
		PlaceholderFormat(sq.Dollar).
		Suffix(`RETURNING "id"`)

	return exec.InsertWithReturningID(ctx, query, db)
}

func (h House) GetByID(ctx context.Context, id uint) (domain.House, error) {
	db := h.queryEngineProvider.GetQueryEngine(ctx)

	query := sq.Select(schema.Flat{}.Columns()...).
		From(houseTable).
		Where("id = $1", id).
		PlaceholderFormat(sq.Dollar)

	house, err := exec.ScanOne[schema.House](ctx, query, db)
	if err != nil {
		return domain.House{}, err
	}

	return schema.NewDomainHouse(house), nil
}

func (h House) UpdateLastFlatAddedAt(ctx context.Context, id uint, lastFlatAddedAt time.Time) error {
	db := h.queryEngineProvider.GetQueryEngine(ctx)

	query := sq.Update(houseTable).
		Set("last_flat_added_at", lastFlatAddedAt).
		Where("id = $1", id).
		PlaceholderFormat(sq.Dollar)

	return exec.Update(ctx, query, db)
}
