package repo

import (
	"backend-bootcamp-assignment-2024/internal/domain"
	"backend-bootcamp-assignment-2024/internal/repo/exec"
	"backend-bootcamp-assignment-2024/internal/repo/schema"
	"backend-bootcamp-assignment-2024/internal/repo/transactor"
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

const (
	flatTable = "bootcamp.flats"
)

type (
	Flat struct {
		queryEngineProvider transactor.QueryEngineProvider
	}
)

func (f Flat) Create(ctx context.Context, flat domain.Flat) (uint, error) {
	db := f.queryEngineProvider.GetQueryEngine(ctx)

	record := schema.NewFlat(flat)
	query := sq.Insert(flatTable).
		Columns(record.Columns()...).
		Values(record.Values()...).
		PlaceholderFormat(sq.Dollar).
		Suffix(`RETURNING "id"`)

	rawQuery, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}

	row, err := db.Query(ctx, rawQuery, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	var id uint
	err = row.Scan(&id)
	return id, err
}

func (f Flat) GetByID(ctx context.Context, id uint) (domain.Flat, error) {
	db := f.queryEngineProvider.GetQueryEngine(ctx)

	query := sq.Select(schema.Flat{}.Columns()...).
		From(flatTable).
		Where("id = $1", id).
		PlaceholderFormat(sq.Dollar)

	flat, err := exec.ScanOne[schema.Flat](ctx, query, db)
	if err != nil {
		return domain.Flat{}, err
	}

	return schema.NewDomainFlat(flat), nil
}

func (f Flat) UpdateStatus(ctx context.Context, id uint, status domain.FlatStatus) error {
	db := f.queryEngineProvider.GetQueryEngine(ctx)

	query := sq.Update(flatTable).
		Set("status", status).
		Where("id = $1", id).
		PlaceholderFormat(sq.Dollar)

	return exec.Update(ctx, query, db)
}

func (f Flat) SetModeratorID(ctx context.Context, id uint, moderatorID *uuid.UUID) error {
	db := f.queryEngineProvider.GetQueryEngine(ctx)

	var nullableModeratorID uuid.UUID
	if moderatorID != nil {
		nullableModeratorID = *moderatorID
	}

	sqlModeratorID := sql.Null[uuid.UUID]{V: nullableModeratorID, Valid: nullableModeratorID.String() != uuid.UUID{}.String()}
	query := sq.Update(flatTable).
		Set("moderator_id", sqlModeratorID).
		Where("id = $1", id).
		PlaceholderFormat(sq.Dollar)

	return exec.Update(ctx, query, db)
}

func NewFlatRepo(provider transactor.QueryEngineProvider) Flat {
	return Flat{provider}
}
