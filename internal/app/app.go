package app

import (
	"context"
	"github.com/core-go/health"
	"github.com/core-go/log"
	"github.com/core-go/search/query"
	q "github.com/core-go/sql"
	"github.com/core-go/sql/client"
	_ "github.com/go-sql-driver/mysql"
	"reflect"

	"go-service/internal/user/adapter/handler"
	userProxy "go-service/internal/user/adapter/proxy"
	"go-service/internal/user/domain"
	"go-service/internal/user/port"
	"go-service/internal/user/service"
)

type ApplicationContext struct {
	Health *health.Handler
	User   port.UserHandler
}

func NewApp(ctx context.Context, conf Config) (*ApplicationContext, error) {
	db, err := q.OpenByConfig(conf.Sql)
	if err != nil {
		return nil, err
	}
	logError := log.ErrorMsg

	userType := reflect.TypeOf(domain.User{})
	userQueryBuilder := query.NewBuilder(db, "users", userType)
	userSearchBuilder, err := q.NewSearchBuilder(db, userType, userQueryBuilder.BuildQuery)
	if err != nil {
		return nil, err
	}

	client0, err := client.NewClient(conf.Client)
	if err != nil {
		return nil, err
	}
	proxy := client.NewProxyClient(client0, conf.ClientProxyUrl, log.InfoFields)
	sqlLayer := userProxy.NewSqlLayerAdapter(proxy, q.BuildParam)
	//userRepository := repository.NewUserAdapter(db)
	userService := service.NewUserService(proxy, sqlLayer)
	userHandler := handler.NewUserHandler(userSearchBuilder.Search, userService, logError)

	sqlChecker := q.NewHealthChecker(db)
	healthHandler := health.NewHandler(sqlChecker)

	return &ApplicationContext{
		Health: healthHandler,
		User:   userHandler,
	}, nil
}
