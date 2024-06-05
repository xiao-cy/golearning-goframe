package auth

import (
	"context"

	v1 "golearning-goframe/api/auth/v1"
	"golearning-goframe/internal/service"
)

func (c *ControllerV1) AdminLogout(ctx context.Context, req *v1.AdminLogoutReq) (res *v1.AdminLogoutRes, err error) {
	service.Auth().LogoutHandler(ctx)
	return
}
