package shorturl

import (
	"context"

	"github.com/colinrs/goshorturl/internal/svc"
	"github.com/colinrs/goshorturl/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateShortUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateShortUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateShortUrlLogic {
	return &CreateShortUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateShortUrlLogic) CreateShortUrl(req *types.CreateShortUrlRequest) (resp *types.CreatShortUrlResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
