package exec

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/repo/transactor"
)

func ScanOne[T any](ctx context.Context, query sq.SelectBuilder, db transactor.QueryEngine) (T, error) {
	var defaultT T

	rawQuery, args, err := query.ToSql()
	if err != nil {
		return defaultT, err
	}

	var records []T
	if err := pgxscan.Select(ctx, db, &records, rawQuery, args...); err != nil {
		return defaultT, err
	}

	return records[0], nil
}
