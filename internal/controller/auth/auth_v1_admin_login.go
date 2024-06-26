package auth

import (
	"context"

	v1 "golearning-goframe/api/auth/v1"
	"golearning-goframe/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) AdminLogin(ctx context.Context, req *v1.AdminLoginReq) (res *v1.AdminLoginRes, err error) {
	res = &v1.AdminLoginRes{}
	res.Token, res.Expire = service.Auth().LoginHandler(ctx)
	g.Log().Line(true).Info(ctx, res)
	return
}
