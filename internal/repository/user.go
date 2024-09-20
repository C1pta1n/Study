package repository

import (
	"context"
	entity "waifunet/internal/entity"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

//TODO: func NewUserRepository
// func ()

// TODO: Прописать сами функции
func (r *UserRepo) Authorize(login, password string) error {
	return nil
}

func (r *UserRepo) GetAll(ctx context.Context) ([]*entity.User, error) {
	return []*entity.User{}, nil
}

func (r *UserRepo) GetManyUsers(ctx context.Context) ([]*entity.User, error) {
	return []*entity.User{}, nil
}

func (r *UserRepo) Create(ctx context.Context, user *entity.User) error {
	return nil
}

func (r *UserRepo) Update(ctx context.Context, user *entity.User) error {
	return nil
}

func (r *UserRepo) Delete(deleteCode string, user *entity.User) error {
	return nil
}
