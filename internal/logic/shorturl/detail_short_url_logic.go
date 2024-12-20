package shorturl

import (
	"context"
	"time"

	"github.com/colinrs/goshorturl/internal/manager"
	"github.com/colinrs/goshorturl/internal/model"
	"github.com/colinrs/goshorturl/internal/svc"
	"github.com/colinrs/goshorturl/internal/types"
	"github.com/colinrs/goshorturl/pkg/code"
	"github.com/colinrs/goshorturl/pkg/gosafe"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type DetailShortUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	shortUrlManager manager.ShortUrlManager
}

func NewDetailShortUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailShortUrlLogic {
	return &DetailShortUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		shortUrlManager: manager.NewShortUrlManager(ctx, svcCtx),
	}
}

func (l *DetailShortUrlLogic) DetailShortUrl(req *types.DetailShortUrlRequest) (resp *types.DetailShortUrlResponse, err error) {
	shortUrl := &model.ShortUrl{
		Model: gorm.Model{
			ID: req.Id,
		},
		ShortUrl: req.Url,
	}
	err = l.shortUrlManager.FindShortUrl(shortUrl)
	if err != nil {
		return
	}
	if shortUrl.ExpireAt.Time.UTC().Unix() < time.Now().UTC().Unix() {
		gosafe.GoSafe(context.WithoutCancel(l.ctx), func() {
			_ = l.svcCtx.LocalCache.Delete(context.WithoutCancel(l.ctx), req.Url)
			err = l.shortUrlManager.DelShortUrl(&model.ShortUrl{ShortUrl: shortUrl.ShortUrl})
			if err != nil {
				logx.Errorf("delete short url err: %v", err)
			}
		})
		return nil, code.UrlNotExist
	}
	resp = &types.DetailShortUrlResponse{
		Id:          shortUrl.ID,
		Url:         shortUrl.ShortUrl,
		Origin:      shortUrl.OriginUrl,
		Description: shortUrl.Description.String,
		ExpireAt:    shortUrl.ExpireAt.Time.String(),
		CreatedAt:   shortUrl.CreatedAt.String(),
	}
	return
}
