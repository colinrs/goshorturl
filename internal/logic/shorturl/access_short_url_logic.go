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
)

type AccessShortUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	shortUrlManager  manager.ShortUrlManager
	accessLogManager manager.AccessLogManager
}

func NewAccessShortUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccessShortUrlLogic {
	return &AccessShortUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		shortUrlManager:  manager.NewShortUrlManager(ctx, svcCtx),
		accessLogManager: manager.NewAccessLogManager(ctx, svcCtx),
	}
}

func (l *AccessShortUrlLogic) AccessShortUrl(req *types.AccessShortUrlRequest) (resp *types.AccessShortUrlResponse, err error) {
	shortUrl := &localShortUrl{}
	err = l.svcCtx.LocalCache.Load(l.ctx, l.loaderShortUrl, req.Url, shortUrl, time.Minute*30)
	if err != nil {
		return nil, err
	}
	if shortUrl.ExpireAt.UTC().Unix() < time.Now().UTC().Unix() {
		gosafe.GoSafe(context.WithoutCancel(l.ctx), func() {
			_ = l.svcCtx.LocalCache.Delete(context.WithoutCancel(l.ctx), req.Url)
			_ = l.shortUrlManager.DelShortUrl(&model.ShortUrl{ShortUrl: shortUrl.ShortUrl})
		})
		return nil, code.UrlNotExist
	}
	resp = &types.AccessShortUrlResponse{
		Localtion: shortUrl.OriginUrl,
	}
	return
}

func (l *AccessShortUrlLogic) loaderShortUrl(ctx context.Context, keys []string) ([]interface{}, error) {
	if len(keys) == 0 {
		return nil, code.UrlNotExist
	}
	url := keys[0]
	shortUrl := &model.ShortUrl{
		ShortUrl: url,
	}
	err := l.shortUrlManager.FindShortUrl(shortUrl)
	if err != nil {
		return nil, err
	}
	var res []interface{}
	res = append(res, localShortUrl{
		ShortUrl:  shortUrl.ShortUrl,
		OriginUrl: shortUrl.OriginUrl,
		ExpireAt:  shortUrl.ExpireAt.Time,
	})
	return res, nil
}

type localShortUrl struct {
	ShortUrl  string    `json:"short_url"`
	OriginUrl string    `json:"origin_url"`
	ExpireAt  time.Time `json:"expire_at"`
}

func (l *AccessShortUrlLogic) SaveAccessLog(accessLog *model.UrlAccessLog) error {
	gosafe.GoSafe(context.WithoutCancel(l.ctx), func() {
		_ = l.accessLogManager.Create(accessLog)
	})
	return nil
}
