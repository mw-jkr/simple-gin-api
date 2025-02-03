package service

import (
	"context"
	"example/test-golang/entity"
	"example/test-golang/repository"
)

type UserService interface {
	GetUserById(ctx context.Context, id int) (*entity.User, error)
}

type userServiceImpl struct {
	repository.UserRepository
}

func (u userServiceImpl) GetUserById(ctx context.Context, id int) (*entity.User, error) {
	user, err := u.UserRepository.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userServiceImpl{UserRepository: ur}
}

func ProvideUserService(ur repository.UserRepository) UserService {
	return NewUserService(ur)
}
