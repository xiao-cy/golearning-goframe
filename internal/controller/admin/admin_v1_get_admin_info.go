package admin

import (
	"context"

	v1 "golearning-goframe/api/admin/v1"
	"golearning-goframe/internal/service"
)

type authController struct{}

var Auth = authController{}

func (c *ControllerV1) GetAdminInfo(ctx context.Context, req *v1.GetAdminInfoReq) (res *v1.GetAdminInfoRes, err error) {

	entityAdmin, err := service.Admin().GetAdminByAdminId(ctx, req.AdminId)

	res = &v1.GetAdminInfoRes{
		AdminInfo: &v1.AdminInfo{
			AdminId:    int64(entityAdmin.AdminId),
			AdminName:  entityAdmin.AdminName,
			Department: entityAdmin.Department,
		},
	}

	return
}
