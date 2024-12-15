package shorturl

import (
	"context"

	"github.com/colinrs/goshorturl/internal/svc"
	"github.com/colinrs/goshorturl/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccessShortUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccessShortUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccessShortUrlLogic {
	return &AccessShortUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccessShortUrlLogic) AccessShortUrl(req *types.AccessShortUrlRequest) (resp *types.AccessShortUrlResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
