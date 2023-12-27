package repository

import (
	"database/sql"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func Timestamptz(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{Time: t, Valid: true}
}

func TimestamptzNullTime(t sql.NullTime) pgtype.Timestamptz {
	return pgtype.Timestamptz{Time: t.Time, Valid: t.Valid}
}
