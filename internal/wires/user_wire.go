//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/google/wire"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/controllers"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/repos"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/services"
)

func InitUserRouterHandler() (*controllers.UserController, error) {
	wire.Build(
		repos.NewUserRepo,
		services.NewUserService,
		controllers.NewUserController,
	)
	return new(controllers.UserController), nil
}
