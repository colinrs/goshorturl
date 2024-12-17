package manager

import (
	"context"

	"github.com/colinrs/goshorturl/internal/model"
	"github.com/colinrs/goshorturl/internal/repo"
	"github.com/colinrs/goshorturl/internal/svc"

	"github.com/spaolacci/murmur3"
	"gorm.io/gorm"
)

type UrlManager interface {
	UrlToShortUrl(url string) string
}

type urlManager struct {
	ctx context.Context
	svc *svc.ServiceContext

	db           *gorm.DB
	shortUrlRepo repo.ShortUrlRepo
}

func NewUrlManager(ctx context.Context, svc *svc.ServiceContext) UrlManager {
	return &urlManager{
		ctx:          ctx,
		svc:          svc,
		db:           svc.GetDB(ctx),
		shortUrlRepo: repo.NewShortUrlRepo(ctx, svc),
	}
}

func (s *urlManager) UrlToShortUrl(url string) string {
	// 需要增加一下判断，如果url已经存在，则直接返回
	shortUrl := &model.ShortUrl{OriginUrl: url}
	err := s.shortUrlRepo.FindShortUrl(s.db, shortUrl)
	if err != nil {
		return ""
	}
	if shortUrl.ShortUrl != "" {
		return shortUrl.ShortUrl
	}
	return base62Encode(murmur3.Sum64([]byte(url)))
}

func base62Encode(id uint64) string {
	const base = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if id == 0 {
		return "0"
	}
	var encoding string
	for id > 0 {
		encoding = string(base[id%62]) + encoding
		id /= 62
	}
	return encoding
}
