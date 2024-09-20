package repository

import (
	"context"
	entity "waifunet/internal/entity"
)

// репозиторий обычного юзера (обращение в бд)
type User interface {
	Authorize(login, password string) error
	GetAll(ctx context.Context) ([]*entity.User, error)
	GetManyUsers(ctx context.Context /*filter*/) ([]*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(deleteCode string, user *entity.User) error
}

// репозиторий админа (обращение в бд)
type Admin interface {
	Authorize(login, password string) error
	GetAll(ctx context.Context) ([]*entity.User, error)
	GetManyUsers(ctx context.Context /*filter*/) ([]*entity.User, error)
	DeleteUser(ctx context.Context, user *entity.User) error
	BlockUser(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
}

type Repository struct {
	User
	// Admin
}

// TODO: добавить функцию создания новых репозиториев
// func NewRepos() {
//
// }
