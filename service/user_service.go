package service

import (
	"errors"
	"example/test-golang/entity"
	"example/test-golang/repository"
)

type UserService interface {
	GetUserById(id int) (*entity.User, error)
}

type userServiceImpl struct {
	repository.UserRepository
}

func (u userServiceImpl) GetUserById(id int) (*entity.User, error) {
	if id == 1 {
		user := &entity.User{Name: "test"}
		return user, nil
	}

	return nil, errors.New("user not found")
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userServiceImpl{UserRepository: ur}
}

func ProvideUserService(ur repository.UserRepository) UserService {
	return NewUserService(ur)
}
