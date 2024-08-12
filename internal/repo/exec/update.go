package exec

import (
	"backend-bootcamp-assignment-2024/internal/repo/repoerr"
	"backend-bootcamp-assignment-2024/internal/repo/transactor"
	"context"
	sq "github.com/Masterminds/squirrel"
)

func Update(ctx context.Context, query sq.UpdateBuilder, db transactor.QueryEngine) error {
	rawQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	tag, err := db.Exec(ctx, rawQuery, args...)
	if err == nil && tag.RowsAffected() == 0 {
		return repoerr.ErrNotFound
	}
	return err
}
