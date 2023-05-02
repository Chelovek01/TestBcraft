package postgresadapter

import (
	"TestBcraft/internal/domain/entity"
	"TestBcraft/pkg/logger"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userStorage struct {
	db *pgxpool.Pool
}

func NewUserStorage(db *pgxpool.Pool) *userStorage {
	return &userStorage{db: db}
}

func (us *userStorage) Create(user *entity.User) error {

	_, err := us.db.Exec(context.Background(), "insert into users (username, password) values ($1, $2)",
		user.Username, user.Password)

	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	return err
}

func (us *userStorage) GetUserByName(username string) (*entity.User, error) {

	var user entity.User

	rows, err := us.db.Query(context.Background(), "select * from users where usename=$1", username)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	if rows.Next() {
		for rows.Next() {
			err := rows.Scan(&user.Id, &user.Username, &user.Password)
			if err != nil {
				logger.ErrorLogger.Println(err)
			}
		}
	}
	return &user, err
}

func (us *userStorage) GetUserById(userId int) (*entity.User, error) {

	var user entity.User

	row, err := us.db.Query(context.Background(), "select * from users where id=$1", userId)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil, err
	}
	if row.Next() {
		for row.Next() {
			err = row.Scan(&user.Id, &user.Username, &user.Password)
			if err != nil {
				logger.ErrorLogger.Println(err)
				return nil, err
			}
		}
	}
	return &user, nil
}

//func (us *userStorage) LoginCheck() error {
//	return nil
//}
