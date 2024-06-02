package admin

import (
	"context"

	v1 "golearning-goframe/api/admin/v1"
	"golearning-goframe/internal/model"
	"golearning-goframe/internal/service"
)

func (c *ControllerV1) CreateAdmin(ctx context.Context, req *v1.CreateAdminReq) (res *v1.CreateAdminRes, err error) {
	res = &v1.CreateAdminRes{}
	var Ainput = model.AdminCreateInput{
		Password:   req.Password,
		AdminName:  req.AdminName,
		Mobile:     req.Mobile,
		Department: req.Department.String(),
	}

	res.AdminId, err = service.Admin().CreateAdmin(ctx, Ainput)

	if err != nil {
		return nil, err
	}
	return
}
