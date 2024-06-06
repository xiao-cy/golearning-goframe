package auth

import (
	"context"

	v1 "golearning-goframe/api/auth/v1"
	"golearning-goframe/internal/service"
)

func (c *ControllerV1) SLogout(ctx context.Context, req *v1.SLogoutReq) (res *v1.SLogoutRes, err error) {
	res = &v1.SLogoutRes{}
	service.Session().RemoveAdmin(ctx)
	return res, nil
}
