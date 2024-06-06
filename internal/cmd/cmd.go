package cmd

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gsession"
	"github.com/gogf/gf/v2/os/gtime"

	"golearning-goframe/internal/controller/admin"
	"golearning-goframe/internal/controller/auth"
	"golearning-goframe/internal/controller/damo"
	"golearning-goframe/internal/controller/hello"
	"golearning-goframe/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 设置session
			s.SetSessionMaxAge(time.Minute)
			s.SetSessionStorage(gsession.NewStorageMemory())
			// 设置路由
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Ctx,
					ghttp.MiddlewareHandlerResponse,
				)
				group.Bind(
					admin.NewV1(),
					auth.NewV1(),
				)
				//session 测试
				group.ALL("/set", func(r *ghttp.Request) {
					r.Session.MustSet("time", gtime.Timestamp())
					r.Response.Write("ok")
				})
				group.ALL("/get", func(r *ghttp.Request) {
					r.Response.Write(r.Session.Data())
				})
				group.ALL("/del", func(r *ghttp.Request) {
					_ = r.Session.RemoveAll()
					r.Response.Write("ok")
				})
				//登录访问功能
				s.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Ctx)
					group.Middleware(
						service.Middleware().Ctx,
						service.Middleware().AdminAuth,
					)
					group.Bind(
						hello.NewV1(),
						damo.NewV1(),
					)
				})
			})

			// s.Group("/", func(group *ghttp.RouterGroup) {
			// 	group.Middleware(ghttp.MiddlewareHandlerResponse)
			// 	group.Bind(
			// 		admin.NewV1(),
			// 		auth.NewV1(),
			// 		hello.NewV1(),
			// 		damo.NewV1(),
			// 	)
			// })

			s.Run()
			return nil
		},
	}
)
