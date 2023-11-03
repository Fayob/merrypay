package port

import (
	"context"
	"merrypay/service"
)

type Server interface {
	CreateUser(context.Context, service.CreateUserParams) (string, error)
	FindUser(context.Context, string) (service.User, error)
	FindAllUsers(context.Context) ([]service.User, error)
	UpdateUser(context.Context, service.UpdateUserParams) (string, error)
	UpdatePassword(context.Context, string, string) (string, error)
	DeleteUser(context.Context, string) (string, error)
}
