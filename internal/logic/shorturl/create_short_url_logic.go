package shorturl

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/colinrs/goshorturl/pkg/gosafe"
	"time"

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
	// 需要增加一下判断，如果url已经存在，则直接返回
	shortUrl := &model.ShortUrl{OriginUrl: req.Origin}
	_ = l.shortUrlManager.FindShortUrl(shortUrl)
	var url string
	var isNew bool
	var id uint
	if shortUrl.ShortUrl != "" {
		url = shortUrl.ShortUrl
		id = shortUrl.ID
	} else {
		isNew = true
		url = l.urlManager.UrlToShortUrl(req.Origin)
	}
	if isNew {
		newShortUrl := &model.ShortUrl{
			OriginUrl: req.Origin,
			ShortUrl:  url,
			Description: sql.NullString{
				String: req.Description,
				Valid:  req.Description != "",
			},
			ExpireAt: sql.NullTime{
				Time:  expireAt,
				Valid: true,
			},
		}
		err = l.shortUrlManager.SaveShortUrl(newShortUrl)
		if err != nil {
			return
		}
		id = newShortUrl.ID
		shortUrl.ShortUrl = newShortUrl.ShortUrl
		shortUrl.OriginUrl = newShortUrl.OriginUrl
		shortUrl.ExpireAt = newShortUrl.ExpireAt
	}
	lc := localShortUrl{
		ShortUrl:  shortUrl.ShortUrl,
		OriginUrl: shortUrl.OriginUrl,
		ExpireAt:  shortUrl.ExpireAt.Time,
	}
	resp = &types.CreatShortUrlResponse{
		Id: id,
		Url: fmt.Sprintf("%s/api/v1/shorturl/access?url=%s",
			l.svcCtx.Config.ShortUrlDomain, url),
	}
	gosafe.GoSafe(context.WithoutCancel(l.ctx), func() {
		err = l.svcCtx.LocalCache.Set(context.WithoutCancel(l.ctx), url, lc, time.Minute*30)
		if err != nil {
			l.Errorf("set url:%s,local cache err:%s", url, err.Error())
		}
		return
	})
	return
}
