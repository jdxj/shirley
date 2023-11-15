package middleware

import "github.com/gogf/gf/v2/net/ghttp"

func CORS(r *ghttp.Request) {
	cors := r.Response.DefaultCORSOptions()
	cors.AllowDomain = []string{"localhost"}
	r.Response.CORS(cors)
	r.Middleware.Next()
}
