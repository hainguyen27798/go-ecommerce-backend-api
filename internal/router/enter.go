package router

import (
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/router/admin"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/router/auth"
	"github.com/hainguyen27798/go-ecommerce-backend-api.git/internal/router/web"
)

type GroupRouter struct {
	Admin admin.GroupRouter
	Auth  auth.RouterGroup
	Web   web.GroupRouter
}

var AppRouter = new(GroupRouter)
