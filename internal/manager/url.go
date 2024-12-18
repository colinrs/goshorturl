package manager

import (
	"context"

	"github.com/colinrs/goshorturl/internal/repo"
	"github.com/colinrs/goshorturl/internal/svc"

	"github.com/spaolacci/murmur3"
	"github.com/teris-io/shortid"
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
	url = base62Encode(murmur3.Sum64([]byte(url)))
	count, err := s.shortUrlRepo.CountShortUrl(s.db, url)
	if err != nil || count > 0 {
		var sid string
		for i := 0; i < 3; i++ {
			sid, err = shortid.Generate()
			if err == nil && sid != "" {
				break
			}
		}
		url = url + sid
	}
	return url
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
