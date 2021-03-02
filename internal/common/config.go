package common

import (
	"github.com/donech/tool/xdb"
	"github.com/donech/tool/xjwt"
	"github.com/donech/tool/xlog"
	"github.com/donech/tool/xredis"
	"github.com/spf13/viper"
)

type Config struct {
	Redis xredis.Config
	DB    xdb.Config
	Jwt   xjwt.Config
}

func InitConfig() Config {
	cfg := Config{}
	err := viper.Unmarshal(cfg)
	if err != nil {
		xlog.SS().Fatalf("init config error")
	}
	return cfg
}
