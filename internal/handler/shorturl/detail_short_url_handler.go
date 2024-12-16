package shorturl

import (
	"net/http"

	"github.com/colinrs/goshorturl/internal/logic/shorturl"
	"github.com/colinrs/goshorturl/internal/svc"
	"github.com/colinrs/goshorturl/internal/types"
	"github.com/colinrs/goshorturl/pkg/code"
	"github.com/colinrs/goshorturl/pkg/httpy"
)

func DetailShortUrlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DetailShortUrlRequest
		if err := httpy.Parse(r, &req); err != nil {
			httpy.ResultCtx(r, w, nil, err)
			return
		}
		if req.Url == "" && req.Id == 0 {
			httpy.ResultCtx(r, w, nil, code.ErrParam)
			return
		}
		l := shorturl.NewDetailShortUrlLogic(r.Context(), svcCtx)
		resp, err := l.DetailShortUrl(&req)
		if err != nil {
			httpy.ResultCtx(r, w, nil, err)
			return
		}
		httpy.ResultCtx(r, w, resp, err)
	}
}
