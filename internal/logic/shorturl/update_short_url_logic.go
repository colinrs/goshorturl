package shorturl

import (
	"context"
	"database/sql"
	"github.com/colinrs/goshorturl/internal/manager"
	"github.com/colinrs/goshorturl/internal/model"
	"github.com/colinrs/goshorturl/internal/svc"
	"github.com/colinrs/goshorturl/internal/types"
	"github.com/colinrs/goshorturl/pkg/utils"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateShortUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	shortUrlManager manager.ShortUrlManager
}

func NewUpdateShortUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateShortUrlLogic {
	return &UpdateShortUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		shortUrlManager: manager.NewShortUrlManager(ctx, svcCtx),
	}
}

func (l *UpdateShortUrlLogic) UpdateShortUrl(req *types.UpdateShortUrlRequest) (resp *types.UpdateShortUrlResponse, err error) {
	if req.Description == "" && req.ExpireAt == "" {
		return nil, nil
	}
	var expireAt time.Time

	if req.ExpireAt != "" {
		expireAt, err = utils.StrTime(req.ExpireAt)
		if err != nil {
			return
		}
	}

	shortUrl := &model.ShortUrl{
		Description: sql.NullString{String: req.Description, Valid: req.Description != ""},
		ExpireAt:    sql.NullTime{Time: expireAt, Valid: !expireAt.IsZero()},
	}
	err = l.shortUrlManager.UpdateShortUrlByID(req.Id, shortUrl)
	if err != nil {
		return
	}
	resp = &types.UpdateShortUrlResponse{}
	return
}
