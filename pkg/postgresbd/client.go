package postgresbd

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

func NewClientPostgres() (*pgxpool.Pool, error) {

	urlExample := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	dbpool, err := pgxpool.New(context.Background(), urlExample)
	if err != nil {
		fmt.Println(err)
	}

	return dbpool, err
}
