package handler

import (
	"github.com/Prakharpandey007/paypocket/internal/service"
)

type Container struct {
	UserHandler *UserHandler
}

func NewContainer(userService service.UserService) *Container {
	return &Container{
		UserHandler: NewUserHandler(userService),
	}
}
