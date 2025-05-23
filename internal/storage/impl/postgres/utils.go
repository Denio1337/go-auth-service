package postgres

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func isUniqueConstraintError(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == "23505" // unique_violation
	}
	return false
}
