package repo

import (
	"context"

	"github.com/colinrs/goshorturl/internal/model"
	"github.com/colinrs/goshorturl/internal/svc"

	"gorm.io/gorm"
)

type AccessLogRepo interface {
	Create(db *gorm.DB, accessLog *model.UrlAccessLog) error
}

type accessLog struct {
	ctx context.Context
	svc *svc.ServiceContext
}

func NewAccessLogRepo(ctx context.Context, svc *svc.ServiceContext) AccessLogRepo {
	return &accessLog{
		ctx: ctx,
		svc: svc,
	}
}

func (a *accessLog) Create(db *gorm.DB, accessLog *model.UrlAccessLog) error {
	return db.Create(accessLog).Error
}
