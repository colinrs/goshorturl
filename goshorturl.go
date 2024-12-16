package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/colinrs/goshorturl/internal/config"
	"github.com/colinrs/goshorturl/internal/handler"
	"github.com/colinrs/goshorturl/internal/svc"
	"github.com/colinrs/goshorturl/pkg/response"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/goshorturl.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/s",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("goshorturl"))
		},
	})
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	httpx.SetErrorHandlerCtx(response.ErrHandle)
	httpx.SetOkHandler(response.OKHandle)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
