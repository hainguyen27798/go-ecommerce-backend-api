package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/dto"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/services"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/pkg/response"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController(userService services.IUserService) *UserController {
	return &UserController{
		userService,
	}
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	response.SuccessResponse(c, response.ErrCodeSuccess, uc.userService.GetUsers())
}

func (uc *UserController) Register(c *gin.Context) {
	var params dto.UserRegistrationRequestDTO
	if err := c.ShouldBindBodyWithJSON(&params); err != nil {
		response.ErrorResponse(c, response.ErrSendEmailFailed, err)
		return
	}
	response.SuccessResponse(c, response.ErrCodeSuccess, uc.userService.Register(params.Email, params.Password))
}
