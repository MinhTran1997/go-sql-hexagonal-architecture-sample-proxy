package service

import (
	"context"
	"github.com/core-go/sql"
	"go-service/internal/user/domain"
	. "go-service/internal/user/port"
)

type UserService interface {
	Load(ctx context.Context, id string) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) (int64, error)
	Update(ctx context.Context, user *domain.User) (int64, error)
	Patch(ctx context.Context, user map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewUserService(proxy sql.Proxy, repository UserRepository) UserService {
	return &userService{
		proxy:      proxy,
		repository: repository,
	}
}

type userService struct {
	proxy      sql.Proxy
	repository UserRepository
}

const (
	timeout = 5000000000
)

func (s *userService) Load(ctx context.Context, id string) (*domain.User, error) {
	return s.repository.Load(ctx, id)
}
func (s *userService) Create(ctx context.Context, user *domain.User) (int64, error) {
	ctx, tx, err := sql.BeginTx(ctx, s.proxy, timeout)
	if err != nil {
		return -1, nil
	}
	res, err := s.repository.Create(ctx, user)
	return sql.EndTx(ctx, s.proxy, tx, res, err)
}
func (s *userService) Update(ctx context.Context, user *domain.User) (int64, error) {
	ctx, tx, err := sql.BeginTx(ctx, s.proxy, timeout)
	if err != nil {
		return -1, nil
	}
	res, err := s.repository.Update(ctx, user)
	return sql.EndTx(ctx, s.proxy, tx, res, err)
}
func (s *userService) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	ctx, tx, err := sql.BeginTx(ctx, s.proxy, timeout)
	if err != nil {
		return -1, nil
	}
	res, err := s.repository.Patch(ctx, user)
	return sql.EndTx(ctx, s.proxy, tx, res, err)
}
func (s *userService) Delete(ctx context.Context, id string) (int64, error) {
	ctx, tx, err := sql.BeginTx(ctx, s.proxy, timeout)
	if err != nil {
		return -1, nil
	}
	res, err := s.repository.Delete(ctx, id)
	return sql.EndTx(ctx, s.proxy, tx, res, err)
}
