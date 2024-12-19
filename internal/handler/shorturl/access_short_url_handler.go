package shorturl

import (
	"database/sql"
	"net/http"

	"github.com/colinrs/goshorturl/internal/logic/shorturl"
	"github.com/colinrs/goshorturl/internal/model"
	"github.com/colinrs/goshorturl/internal/svc"
	"github.com/colinrs/goshorturl/internal/types"
	"github.com/colinrs/goshorturl/pkg/httpy"
)

func AccessShortUrlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AccessShortUrlRequest
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
		// 设置 302 重定向和 Location 头
		w.Header().Set("Location", resp.Localtion)
		w.WriteHeader(http.StatusFound) // 302 状态码
		accessLog := &model.UrlAccessLog{
			ShortUrl:  req.Url,
			Ip:        sql.NullString{String: r.RemoteAddr, Valid: true},
			UserAgent: sql.NullString{String: r.UserAgent(), Valid: true},
			Referrer:  sql.NullString{String: r.Referer(), Valid: true},
		}
		_ = l.SaveAccessLog(accessLog)
	}
}
