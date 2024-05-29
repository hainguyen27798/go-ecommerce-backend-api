package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/services"
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
	c.JSON(200, gin.H{
		"message":   "Success",
		"usernames": uc.userService.GetUsers(),
	})
}
