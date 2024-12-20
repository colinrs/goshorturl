package repo

import (
	"context"
	"github.com/colinrs/goshorturl/internal/model"
	"github.com/colinrs/goshorturl/internal/svc"
	"gorm.io/gorm"
)

type ShortUrlRepo interface {
	SaveShortUrl(db *gorm.DB, shortUrl *model.ShortUrl) error
	FindShortUrl(db *gorm.DB, shortUrl *model.ShortUrl) error
	CountShortUrl(db *gorm.DB, url string) (int64, error)
	UpdateShortUrlByID(db *gorm.DB, id uint, shortUrl *model.ShortUrl) error
	DelShortUrl(db *gorm.DB, shortUrl *model.ShortUrl) error
}

type shortUrlRepo struct {
	ctx context.Context
	svc *svc.ServiceContext
}

func NewShortUrlRepo(ctx context.Context, svc *svc.ServiceContext) ShortUrlRepo {
	return &shortUrlRepo{
		ctx: ctx,
		svc: svc,
	}
}

func (s *shortUrlRepo) SaveShortUrl(db *gorm.DB, shortUrl *model.ShortUrl) error {
	return db.WithContext(s.ctx).Save(shortUrl).Error
}

func (s *shortUrlRepo) FindShortUrl(db *gorm.DB, shortUrl *model.ShortUrl) error {
	return db.WithContext(s.ctx).Where(shortUrl).First(shortUrl).Error
}

func (s *shortUrlRepo) UpdateShortUrlByID(db *gorm.DB, id uint, shortUrl *model.ShortUrl) error {
	return db.WithContext(s.ctx).Model(&model.ShortUrl{}).Where("id = ?", id).Updates(shortUrl).Error
}

func (s *shortUrlRepo) CountShortUrl(db *gorm.DB, url string) (int64, error) {
	var count int64
	err := db.Model(&model.ShortUrl{}).Where("short_url = ?", url).Count(&count).Error
	return count, err

}

func (s *shortUrlRepo) DelShortUrl(db *gorm.DB, shortUrl *model.ShortUrl) error {
	return db.WithContext(context.WithoutCancel(s.ctx)).Where(shortUrl).Delete(&model.ShortUrl{}).Error
}
