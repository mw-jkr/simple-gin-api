package service

import (
	"context"
	"example/test-golang/entity"
)

type ItemService interface {
	GetAll(ctx context.Context) []entity.Item
}

type itemServiceImpl struct {
}

func (i *itemServiceImpl) GetAll(ctx context.Context) []entity.Item {
	return []entity.Item{}
}

func NewItemService() ItemService {
	return &itemServiceImpl{}
}

func ProvideItemService() ItemService {
	return NewItemService()
}
