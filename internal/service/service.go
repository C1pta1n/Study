package service

import (
	"context"
	"waifunet/internal/entity"
	"waifunet/internal/repository"
)

// функционал обычного юзера (фичи)
type User interface {
	SignIn(login, password string) error
	GetAll(ctx context.Context) ([]*entity.User, error)
	GetManyUsers(ctx context.Context /*filter*/) ([]*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(deleteCode string, user *entity.User) error
}

// функционал админа (фичи)
type Admin interface {
	Authorize(login, password string) error
	GetAll(ctx context.Context) ([]*entity.User, error)
	GetManyUsers(ctx context.Context /*filter*/) ([]*entity.User, error)
	DeleteUser(ctx context.Context, user *entity.User) error
	BlockUser(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
}

type Services struct {
	User
	// Admin
}

type Dependencies struct {
	Repos *repository.Repository
}

// TODO: добавить функцию создания новых сервисов
func NewServices(deps Dependencies) *Services {
	return &Services{
		User: NewUserService(deps.Repos.User),
		// Admin: NewAdminUserService()}
	}
}
