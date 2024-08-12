package exec

import (
	"backend-bootcamp-assignment-2024/internal/repo/repoerr"
	"backend-bootcamp-assignment-2024/internal/repo/transactor"
	"context"
	sq "github.com/Masterminds/squirrel"
)

func InsertWithReturningID(ctx context.Context, query sq.InsertBuilder, db transactor.QueryEngine) (uint, error) {
	rawQuery, args, err := query.ToSql()
	if err != nil {
		return 0, err
	}

	row := db.QueryRow(ctx, rawQuery, args...)

	var id uint
	err = row.Scan(&id)

	if err != nil && isDuplicateKeyError(err) {
		return 0, repoerr.ErrDuplicate
	}

	return id, err
}

func Insert(ctx context.Context, query sq.InsertBuilder, db transactor.QueryEngine) error {
	rawQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = db.Exec(ctx, rawQuery, args...)
	if err == nil {
		return nil
	}

	if isDuplicateKeyError(err) {
		return repoerr.ErrDuplicate
	}
	return err
}
