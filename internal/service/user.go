package service

import (
	"context"
	"waifunet/internal/entity"
	"waifunet/internal/repository"
)

type UserService struct {
	repo repository.User
}

// TODO: Прописать сами функции
func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) SignIn(login, password string) error {
	return nil
}
func (s *UserService) GetAll(ctx context.Context) ([]*entity.User, error) {
	return []*entity.User{}, nil
}
func (s *UserService) GetManyUsers(ctx context.Context /*filter*/) ([]*entity.User, error) {
	return []*entity.User{}, nil
}
func (s *UserService) Create(ctx context.Context, user *entity.User) error {
	return nil
}
func (s *UserService) Update(ctx context.Context, user *entity.User) error {
	return nil
}
func (s *UserService) Delete(deleteCode string, user *entity.User) error {
	return nil
}
