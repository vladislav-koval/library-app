package helpers

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

const (
	UniqueViolationSQLState = "23505"
)

func IsUniqueViolation(
	err error,
	constraint string,
) bool {
	pgErr, ok := errors.AsType[*pgconn.PgError](err)

	return ok &&
		pgErr.Code == UniqueViolationSQLState &&
		pgErr.ConstraintName == constraint
}
