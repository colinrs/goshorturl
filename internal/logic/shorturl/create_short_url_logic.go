package shorturl

import (
	"context"
	"database/sql"
	"github.com/colinrs/goshorturl/internal/manager"
	"github.com/colinrs/goshorturl/internal/model"
	"github.com/colinrs/goshorturl/internal/svc"
	"github.com/colinrs/goshorturl/internal/types"
	"github.com/colinrs/goshorturl/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateShortUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	urlManager      manager.UrlManager
	shortUrlManager manager.ShortUrlManager
}

func NewCreateShortUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateShortUrlLogic {
	return &CreateShortUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		urlManager:      manager.NewUrlManager(ctx, svcCtx),
		shortUrlManager: manager.NewShortUrlManager(ctx, svcCtx),
	}
}

func (l *CreateShortUrlLogic) CreateShortUrl(req *types.CreateShortUrlRequest) (resp *types.CreatShortUrlResponse, err error) {

	expireAt, err := utils.StrTime(req.ExpireAt)
	if err != nil {
		return
	}
	shortUrl := &model.ShortUrl{
		OriginUrl: req.Origin,
		ShortUrl:  l.urlManager.UrlToShortUrl(req.Origin),
		Description: sql.NullString{
			String: req.Description,
			Valid:  req.Description != "",
		},
		ExpireAt: sql.NullTime{
			Time:  expireAt,
			Valid: true,
		},
	}
	err = l.shortUrlManager.SaveShortUrl(shortUrl)
	if err != nil {
		return
	}
	resp = &types.CreatShortUrlResponse{
		Id:  shortUrl.ID,
		Url: shortUrl.ShortUrl,
	}
	return
}
