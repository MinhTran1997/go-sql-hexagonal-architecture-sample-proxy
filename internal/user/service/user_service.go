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
	tx, err := s.proxy.BeginTransaction(ctx, timeout)
	if err != nil {
		return -1, nil
	}
	ctx = context.WithValue(ctx, "txId", &tx)
	res, err := s.repository.Create(ctx, user)
	if err != nil {
		err := s.proxy.RollbackTransaction(ctx, tx)
		if err != nil {
			return -1, err
		}
	}
	if err = s.proxy.CommitTransaction(ctx, tx); err != nil {
		return -1, err
	}
	return res, err
	//client := &http.Client{}
	//url := fmt.Sprintf("http://localhost:8080/end?tx=%s", tx)
	//req, err := http.NewRequest("POST", url, nil)
	//if err != nil {
	//	return -1, err
	//}
	//req.Header.Add("Accept", "application/json")
	//req.Header.Add("Content-Type", "application/json")
	//resp, err := client.Do(req)
	//if err != nil {
	//	return -1, err
	//}
	//defer resp.Body.Close()
	//bodyBytes, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return -1, err
	//}
	//var responseObject interface{}
	//json.Unmarshal(bodyBytes, &responseObject)
	//if responseObject == true {
	//	return res, err
	//}
	//
	//return -1, nil
}
func (s *userService) Update(ctx context.Context, user *domain.User) (int64, error) {
	tx, err := s.proxy.BeginTransaction(ctx, timeout)
	if err != nil {
		return -1, nil
	}
	ctx = context.WithValue(ctx, "txId", &tx)
	res, err := s.repository.Update(ctx, user)
	if err != nil {
		err := s.proxy.RollbackTransaction(ctx, tx)
		if err != nil {
			return -1, err
		}
	}
	if err = s.proxy.CommitTransaction(ctx, tx); err != nil {
		return -1, err
	}
	return res, err
}
func (s *userService) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	tx, err := s.proxy.BeginTransaction(ctx, timeout)
	if err != nil {
		return -1, nil
	}
	ctx = context.WithValue(ctx, "txId", &tx)
	res, err := s.repository.Patch(ctx, user)
	if err != nil {
		err := s.proxy.RollbackTransaction(ctx, tx)
		if err != nil {
			return -1, err
		}
	}
	if err = s.proxy.CommitTransaction(ctx, tx); err != nil {
		return -1, err
	}
	return res, nil
}
func (s *userService) Delete(ctx context.Context, id string) (int64, error) {
	tx, err := s.proxy.BeginTransaction(ctx, timeout)
	if err != nil {
		return -1, nil
	}
	ctx = context.WithValue(ctx, "txId", &tx)
	res, err := s.repository.Delete(ctx, id)
	if err != nil {
		err := s.proxy.RollbackTransaction(ctx, tx)
		if err != nil {
			return -1, err
		}
	}
	if err = s.proxy.CommitTransaction(ctx, tx); err != nil {
		return -1, err
	}
	return res, nil
}
