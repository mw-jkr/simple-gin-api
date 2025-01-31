package controller

import (
	"example/test-golang/common"
	"example/test-golang/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUserById(ctx *gin.Context)
}

type userControllerImpl struct {
	service.UserService
}

func (u *userControllerImpl) GetUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.BuildErrorResponse(err))
		return
	}

	user, err := u.UserService.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, common.BuildErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func NewUserController(us service.UserService) UserController {
	return &userControllerImpl{UserService: us}
}

func ProvideUserController(us service.UserService) UserController {
	return NewUserController(us)
}
