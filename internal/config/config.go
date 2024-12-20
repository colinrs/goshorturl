package config

import (
	"github.com/colinrs/goshorturl/internal/infra"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	DataBase       *infra.DBConfig `json:"Database" yaml:"Database"`
	Redis          redis.RedisConf `json:"Redis" yaml:"Redis"`
	ShortUrlDomain string          `json:"ShortUrlDomain" yaml:"ShortUrlDomain"`
}
