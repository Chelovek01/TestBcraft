package postgresadapter

import (
	"TestBcraft/internal/domain/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userStorage struct {
	db *pgxpool.Pool
}

func NewUserStorage(db *pgxpool.Pool) *userStorage {
	return &userStorage{db: db}
}

func (us *userStorage) Create(user *entity.User) error {
	return nil
}
