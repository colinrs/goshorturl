package shorturl

import (
	"context"
	"github.com/colinrs/goshorturl/internal/manager"
	"github.com/colinrs/goshorturl/internal/model"
	"gorm.io/gorm"

	"github.com/colinrs/goshorturl/internal/svc"
	"github.com/colinrs/goshorturl/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
