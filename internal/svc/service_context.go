package svc

import (
	"context"

	"github.com/colinrs/goshorturl/internal/config"
	"github.com/colinrs/goshorturl/internal/infra"
	"github.com/colinrs/goshorturl/pkg/cache"
	"github.com/colinrs/goshorturl/pkg/codec"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config      config.Config
	DB          *gorm.DB
	RedisClient *redis.Redis
	LocalCache  cache.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	s := &ServiceContext{
		Config:      c,
		DB:          initDB(c),
		RedisClient: initRedis(c),
	}
	s.InitLocalCache()
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

func (s *ServiceContext) InitLocalCache() {
	memCache, err := newLocalCache()
	logx.Must(err)
	s.LocalCache = memCache
}

func newLocalCache() (cache.Cache, error) {
	memCache, err := cache.NewRistrettoCache(cache.RistrettoCacheConfig{
		Capacity:    2147483648, // bytes， max mem:2G
		NumCounters: 200000000,  // max keys
		CostFunc:    cache.CostMemoryUsage,
	}, codec.NewSonicCodec())
	if err != nil {
		return nil, err
	}
	return memCache, nil
}
