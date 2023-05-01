package service

import "TestBcraft/internal/domain/entity"

type UserStorage interface {
	Create(user *entity.User) error
}

type userService struct {
	storage UserStorage
}

func NewUserService(storage UserStorage) *userService {
	return &userService{storage: storage}
}

func (us *userService) CreateUser(user *entity.User) error {
	return nil
}
