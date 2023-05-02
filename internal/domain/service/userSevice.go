package service

import (
	"TestBcraft/internal/domain/entity"
	"TestBcraft/internal/token"
	"TestBcraft/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
)

type UserStorage interface {
	Create(user *entity.User) error
	GetUserByName(username string) (*entity.User, error)
	GetUserById(userId int) (*entity.User, error)
}

type userService struct {
	storage UserStorage
}

func NewUserService(storage UserStorage) *userService {
	return &userService{storage: storage}
}

func (us *userService) CreateUser(user *entity.User) error {

	err := us.storage.Create(user)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	return nil
}

func (us *userService) HashPass(user *entity.User) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
	user.Password = string(hashedPassword)

	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	return err
}

func (us *userService) LoginCheck(userLogin *entity.User) (string, error) {

	savedUser, err := us.storage.GetUserByName(userLogin.Username)

	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(savedUser.Password), []byte(userLogin.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		logger.ErrorLogger.Println(err)
		return "", err
	}

	token, errToken := token.GenerateToken(userLogin.Id)
	if errToken != nil {
		logger.ErrorLogger.Println(err)
		return "", err
	}

	return token, nil
}

func (us *userService) GetUserById(userId int) (*entity.User, error) {

	user, err := us.storage.GetUserById(userId)
	if err != nil {
		logger.ErrorLogger.Println(err)
	}

	return user, err
}
