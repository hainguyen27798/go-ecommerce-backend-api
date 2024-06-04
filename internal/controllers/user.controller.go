package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/services"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/pkg/response"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	response.SuccessResponse(c, response.ErrCodeSuccess, uc.userService.GetUsers())
}
