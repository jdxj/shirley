package middleware

import (
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/jdxj/shirley/internal/config"
	"github.com/jdxj/shirley/utility"
)

func Auth(r *ghttp.Request) {
	token := r.GetHeader("token")
	_, err := utility.ParseToken(config.Shirley.Key, token)
	if err != nil {
		g.Log().Errorf(r.GetCtx(), "ParseToken err: %s", err)
		r.Response.WriteStatus(http.StatusForbidden)
		return
	}
	r.Middleware.Next()
}
