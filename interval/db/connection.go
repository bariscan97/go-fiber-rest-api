package db

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"context"
	"os"
)

func Pool() *pgxpool.Pool{
	databaseUrI := os.Getenv("DB_URI")
	dbPool, err := pgxpool.Connect(context.Background(), databaseUrI)
	if err != nil {
		panic(err)
	}
	return dbPool
}