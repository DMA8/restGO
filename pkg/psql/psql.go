package psql

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func ConnectToDB(dsn string) (*pgxpool.Pool, error) {
	db, err := pgxpool.Connect(context.Background(), dsn)
	return db, err
}
