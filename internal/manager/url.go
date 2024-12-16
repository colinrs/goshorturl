package manager

import (
	"context"
	"fmt"
	"github.com/colinrs/goshorturl/internal/svc"
	"github.com/spaolacci/murmur3"
)

type UrlManager interface {
	UrlToShortUrl(url string) string
}

type urlManager struct {
	ctx context.Context
	svc *svc.ServiceContext
}

func NewUrlManager(ctx context.Context, svc *svc.ServiceContext) UrlManager {
	return &urlManager{
		ctx: ctx,
		svc: svc,
	}
}

func (s *urlManager) UrlToShortUrl(url string) string {
	return fmt.Sprintf("%s/api/v1/shorturl/access/%s",
		s.svc.Config.ShortUrlDomain, base62Encode(murmur3.Sum64([]byte(url))))
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
