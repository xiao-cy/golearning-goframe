package middleware

import (
	"golearning-goframe/internal/model"
	"golearning-goframe/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sMiddleware struct {
	AdminLoginUrl string // 登录路由地址
}

func init() {
	service.RegisterMiddleware(New())
}

func New() service.IMiddleware {
	return &sMiddleware{
		AdminLoginUrl: "/admin/login",
	}
}

func (s *sMiddleware) AdminAuth(r *ghttp.Request) {
	g.Log().Line(true).Print(r.Context(), "AdminAuth")
	admin := service.Session().GetAdmin(r.Context())
	if admin.AdminId == 0 {
		// r.Response.RedirectTo(s.AdminLoginUrl)
		r.Response.WriteExit("请先登录！！")
	}
	r.Middleware.Next()
}

func (s *sMiddleware) Ctx(r *ghttp.Request) {
	g.Log().Line(true).Print(r.Context(), "Ctx")
	customCtx := &model.Context{
		Session: r.Session,
	}
	g.Log().Line(true).Info(r.Context(), customCtx)
	service.BizCtx().Init(r, customCtx)
	// 将自定义的上下文对象传递到模板变量中使用
	r.Assigns(g.Map{
		"Context": customCtx,
	})
	r.Middleware.Next()
}
