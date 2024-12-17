package shorturl

import (
	"context"

	"github.com/colinrs/goshorturl/internal/manager"
	"github.com/colinrs/goshorturl/internal/model"
	"github.com/colinrs/goshorturl/internal/svc"
	"github.com/colinrs/goshorturl/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccessShortUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	shortUrlManager manager.ShortUrlManager
}

func NewAccessShortUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccessShortUrlLogic {
	return &AccessShortUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		shortUrlManager: manager.NewShortUrlManager(ctx, svcCtx),
	}
}

func (l *AccessShortUrlLogic) AccessShortUrl(req *types.AccessShortUrlRequest) (resp *types.AccessShortUrlResponse, err error) {
	shortUrl := &model.ShortUrl{
		ShortUrl: req.Url,
	}
	err = l.shortUrlManager.FindShortUrl(shortUrl)
	if err != nil {
		return
	}
	resp = &types.AccessShortUrlResponse{
		Localtion: shortUrl.OriginUrl,
	}
	return
}
