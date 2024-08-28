package services

import (
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/repos"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/pkg/response"
)

type IUserService interface {
	Register(email string, password string) int
	GetUsers() []string
}

type userService struct {
	userRepo repos.IUserRepo
}

func (us userService) GetUsers() []string {
	return us.userRepo.GetUsers()
}

func (us userService) Register(email string, password string) int {
	if us.userRepo.CheckUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}
	return response.ErrCodeSuccess
}

func NewUserService(userRepo repos.IUserRepo) IUserService {
	return &userService{
		userRepo,
	}
}
