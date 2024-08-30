package services

import (
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/repos"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/utils"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/pkg/response"
	"strconv"
	"time"
)

type IUserService interface {
	Register(email string, password string) int
	GetUsers() []string
}

type userService struct {
	userRepo     repos.IUserRepo
	userAuthRepo repos.IUserAuthRepo
}

func NewUserService(userRepo repos.IUserRepo, userAuthRepo repos.IUserAuthRepo) IUserService {
	return &userService{
		userRepo,
		userAuthRepo,
	}
}

func (us userService) GetUsers() []string {
	return us.userRepo.GetUsers()
}

func (us userService) Register(email string, password string) int {
	// hash email
	hashEmail := utils.GetHash(email)

	// check email already exists
	if us.userRepo.CheckUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}

	// new OTP
	otp := utils.GenerateOTO()

	if err := us.userAuthRepo.AddOTP(hashEmail, otp, int64(10*time.Minute)); err != nil {
		return response.ErrInvalidOTP
	}

	if err := utils.SendToEmail(
		"otp-email",
		"hainguyen27798@gmail.com",
		[]string{email},
		map[string]interface{}{
			"otp": strconv.Itoa(otp),
		},
	); err != nil {
		return response.ErrSendEmailFailed
	}

	return response.ErrCodeSuccess
}
