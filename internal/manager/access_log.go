package manager

import (
	"context"
	"github.com/colinrs/goshorturl/internal/model"
	"github.com/colinrs/goshorturl/internal/repo"
	"github.com/colinrs/goshorturl/internal/svc"
	"github.com/colinrs/goshorturl/pkg/code"
	"github.com/colinrs/goshorturl/pkg/gosafe"
	"github.com/zeromicro/go-zero/core/logx"
	"time"

	"gorm.io/gorm"
)

type AccessLogManager interface {
	Create(accessLog *model.UrlAccessLog) error
}

func NewAccessLogManager(ctx context.Context, svc *svc.ServiceContext) AccessLogManager {
	a := &accessLogManager{
		ctx:           ctx,
		svc:           svc,
		db:            svc.GetDB(context.WithoutCancel(ctx)),
		accessLogRepo: repo.NewAccessLogRepo(ctx, svc),
		logChanel:     make(chan *model.UrlAccessLog, 10000),
	}
	gosafe.GoSafe(context.Background(), a.run)
	return a
}

type accessLogManager struct {
	ctx context.Context
	svc *svc.ServiceContext

	db            *gorm.DB
	accessLogRepo repo.AccessLogRepo
	logChanel     chan *model.UrlAccessLog
}

func (a *accessLogManager) Create(accessLog *model.UrlAccessLog) error {
	t := time.Tick(500 * time.Millisecond)
	for {
		select {
		case <-t:
			return code.ErrDatabase
		case a.logChanel <- accessLog:
			return nil
		}
	}
}

func (a *accessLogManager) run() {
	for log := range a.logChanel {
		err := a.accessLogRepo.Create(a.db, log)
		if err != nil {
			logx.Errorf("log:%s,accessLogRepo.Create error: %v", log.ShortUrl, err)
		}
	}
}
