package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"golearning-goframe/api/auth/v1"
)

func (c *ControllerV1) AdminLogout(ctx context.Context, req *v1.AdminLogoutReq) (res *v1.AdminLogoutRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
