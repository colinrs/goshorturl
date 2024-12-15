package shorturl

import (
	"context"

	"github.com/colinrs/goshorturl/internal/svc"
	"github.com/colinrs/goshorturl/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailShortUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailShortUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailShortUrlLogic {
	return &DetailShortUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailShortUrlLogic) DetailShortUrl(req *types.DetailShortUrlRequest) (resp *types.DetailShortUrlResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
