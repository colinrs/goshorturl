package manager

import (
	"context"

	"github.com/colinrs/goshorturl/internal/model"
	"github.com/colinrs/goshorturl/internal/repo"
	"github.com/colinrs/goshorturl/internal/svc"

	"gorm.io/gorm"
)

type ShortUrlManager interface {
	SaveShortUrl(shortUrl *model.ShortUrl) error
	FindShortUrl(shortUrl *model.ShortUrl) error

	UpdateShortUrlByID(id uint, shortUrl *model.ShortUrl) error
}

type shortUrlManager struct {
	ctx context.Context
	svc *svc.ServiceContext

	db           *gorm.DB
	shortUrlRepo repo.ShortUrlRepo
}

func NewShortUrlManager(ctx context.Context, svc *svc.ServiceContext) ShortUrlManager {
	return &shortUrlManager{
		ctx:          ctx,
		svc:          svc,
		db:           svc.DB,
		shortUrlRepo: repo.NewShortUrlRepo(ctx, svc),
	}
}

func (s *shortUrlManager) SaveShortUrl(shortUrl *model.ShortUrl) error {
	return s.shortUrlRepo.SaveShortUrl(s.db, shortUrl)
}

func (s *shortUrlManager) FindShortUrl(shortUrl *model.ShortUrl) error {
	return s.shortUrlRepo.FindShortUrl(s.db, shortUrl)
}

func (s *shortUrlManager) UpdateShortUrlByID(id uint, shortUrl *model.ShortUrl) error {
	return s.shortUrlRepo.UpdateShortUrlByID(s.db, id, shortUrl)
}
