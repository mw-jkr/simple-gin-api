package repository

import (
	"errors"
	"example/test-golang/entity"
)

type UserRepository interface {
	GetUserById(id int) (*entity.User, error)
}

type userRepositoryImpl struct {
}

func (u userRepositoryImpl) GetUserById(id int) (*entity.User, error) {
	if id == 1 {
		user := &entity.User{Name: "test"}
		return user, nil
	}

	return nil, errors.New("user not found")
}

func NewUserRepository() UserRepository {
	return &userRepositoryImpl{}
}

func ProvideUserRepository() UserRepository {
	return NewUserRepository()
}
