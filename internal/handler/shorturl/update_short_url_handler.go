package shorturl

import (
	"net/http"

	"github.com/colinrs/goshorturl/internal/logic/shorturl"
	"github.com/colinrs/goshorturl/internal/svc"
	"github.com/colinrs/goshorturl/internal/types"
	"github.com/colinrs/goshorturl/pkg/httpy"
)

func UpdateShortUrlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateShortUrlRequest
		if err := httpy.Parse(r, &req); err != nil {
			httpy.ResultCtx(r, w, nil, err)
			return
		}

		l := shorturl.NewUpdateShortUrlLogic(r.Context(), svcCtx)
		resp, err := l.UpdateShortUrl(&req)
		if err != nil {
			httpy.ResultCtx(r, w, nil, err)
			return
		}
		httpy.ResultCtx(r, w, resp, err)
	}
}
