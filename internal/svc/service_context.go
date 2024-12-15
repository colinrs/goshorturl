package svc

import (
	"context"

	"github.com/colinrs/goshorturl/internal/config"
	"github.com/colinrs/goshorturl/internal/infra"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config      config.Config
	DB          *gorm.DB
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	s := &ServiceContext{
		Config:      c,
		DB:          initDB(c),
		RedisClient: initRedis(c),
	}
	return s
}

func initDB(c config.Config) *gorm.DB {
	db, err := infra.Database(c.DataBase)
	logx.Must(err)
	return db
}

func initRedis(c config.Config) *redis.Redis {
	return redis.MustNewRedis(c.Redis)
}

func (s *ServiceContext) GetDB(ctx context.Context) *gorm.DB {
	return s.DB.WithContext(ctx)

}
