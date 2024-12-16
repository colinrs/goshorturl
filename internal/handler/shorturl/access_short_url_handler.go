package shorturl

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"github.com/colinrs/goshorturl/internal/logic/shorturl"
	"github.com/colinrs/goshorturl/internal/svc"
	"github.com/colinrs/goshorturl/internal/types"
	"github.com/colinrs/goshorturl/pkg/httpy"
)

func AccessShortUrlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AccessShortUrlRequest
		logx.WithContext(r.Context()).Infof("aaaaa %+v", r.RequestURI)
		if err := httpy.Parse(r, &req); err != nil {
			httpy.ResultCtx(r, w, nil, err)
			return
		}
		l := shorturl.NewAccessShortUrlLogic(r.Context(), svcCtx)
		resp, err := l.AccessShortUrl(&req)
		if err != nil {
			httpy.ResultCtx(r, w, nil, err)
			return
		}
		httpy.ResultCtx(r, w, resp, err)
	}
}
