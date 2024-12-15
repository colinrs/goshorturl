package shorturl

import (
	"context"

	"github.com/colinrs/goshorturl/internal/svc"
	"github.com/colinrs/goshorturl/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateShortUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateShortUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateShortUrlLogic {
	return &UpdateShortUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateShortUrlLogic) UpdateShortUrl(req *types.UpdateShortUrlRequest) (resp *types.UpdateShortUrlResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
