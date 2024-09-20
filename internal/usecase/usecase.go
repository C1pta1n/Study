package usecase

import (
	"context"
	"waifunet/internal/entity"
)

type userService interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUser(ctx context.Context, id int) (entity.User, error)
	GetManyUsers(ctx context.Context) ([]entity.User, error)
}

type UseCase struct {
	userSrv userService
}
