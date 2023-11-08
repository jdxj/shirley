package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"github.com/jdxj/shirley/internal/controller/shares"
	"github.com/jdxj/shirley/internal/controller/token"
	"github.com/jdxj/shirley/internal/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(ghttp.MiddlewareHandlerResponse)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Bind(
					token.NewV1(),
				)
			})
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.Auth)
				group.Bind(
					shares.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
