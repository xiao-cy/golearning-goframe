package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"golearning-goframe/api/auth/v1"
)

func (c *ControllerV1) AdminRefreshToken(ctx context.Context, req *v1.AdminRefreshTokenReq) (res *v1.AdminRefreshTokenRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
