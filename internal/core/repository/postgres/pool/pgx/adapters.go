package core_pgx_pool

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	core_postrgres_pool "github.com/maksimablavatskii/Golang-ToDoap/internal/core/repository/postgres/pool"
)

type pgxRows struct {
	pgx.Rows
}

type pgxRow struct {
	pgx.Row
}

func (r pgxRow) Scan(dest ...any) error {
	err := r.Row.Scan(dest...)
	if err != nil {
		return mapErrors(err)
	}
	return nil
}

type pgxCommandTag struct {
	pgconn.CommandTag
}

func mapErrors(err error) error {
	const (
		pgxViolateForeignKeyErrorCode = "23503"
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return core_postrgres_pool.ErrNoRows
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == pgxViolateForeignKeyErrorCode {
			return fmt.Errorf("%v: %w", err, core_postrgres_pool.ErrViolatesForeignKey)
		}
	}

	return fmt.Errorf("%v: %w", err, core_postrgres_pool.ErrUnKnown)
}
