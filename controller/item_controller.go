package controller

import (
	"example/test-golang/entity/dto"
	"example/test-golang/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemController interface {
	GetAll(ctx *gin.Context)
}

type itemControllerImpl struct {
	service.ItemService
}

func (c *itemControllerImpl) GetAll(ctx *gin.Context) {
	items := c.ItemService.GetAll(ctx)

	dto := dto.GetAllResponse{Items: items}

	ctx.JSON(http.StatusOK, dto)
}

func NewItemController(is service.ItemService) ItemController {
	return &itemControllerImpl{ItemService: is}
}

func ProvideItemController(is service.ItemService) ItemController {
	return NewItemController(is)
}
