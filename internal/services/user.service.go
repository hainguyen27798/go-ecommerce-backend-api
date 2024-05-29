package services

import "github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/repos"

type UserService struct {
	userRepo *repos.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repos.NewUserRepo(),
	}
}

func (us *UserService) GetUsers() []string {
	return us.userRepo.GetUsers()
}
