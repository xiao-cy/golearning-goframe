package auth

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "golearning-goframe/api/auth/v1"
	"golearning-goframe/internal/service"
)

func (c *ControllerV1) SLogin(ctx context.Context, req *v1.SLoginReq) (res *v1.SLoginRes, err error) {
	res = &v1.SLoginRes{}
	admin, _ := service.Admin().GetAdminByMobileAndPassword(ctx, req.Mobile, req.Password)

	if admin != nil {
		g.Log().Line(true).Info(ctx, admin)
		service.Session().SetAdmin(ctx, admin)
		res.AdminId = int(admin.AdminId)
	}

	return res, nil
}
