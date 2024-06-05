package auth

import (
	"context"

	v1 "golearning-goframe/api/auth/v1"
	"golearning-goframe/internal/service"
)

func (c *ControllerV1) AdminRefreshToken(ctx context.Context, req *v1.AdminRefreshTokenReq) (res *v1.AdminRefreshTokenRes, err error) {
	res = &v1.AdminRefreshTokenRes{}
	res.Token, res.Expire = service.Auth().RefreshHandler(ctx)
	return
}
