package app

import (
	"github.com/core-go/log"
	mid "github.com/core-go/log/middleware"
	sv "github.com/core-go/service"
	"github.com/core-go/sql"
	"github.com/core-go/sql/client"
)

type Config struct {
	Server         sv.ServerConf `mapstructure:"server"`
	Sql            sql.Config    `mapstructure:"sql"`
	Client         client.Config `mapstructure:"client"`
	Log            log.Config    `mapstructure:"log"`
	MiddleWare     mid.LogConfig `mapstructure:"middleware"`
	ClientProxyUrl string        `mapstructure:"server_url"`
}
