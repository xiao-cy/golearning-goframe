// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"context"

	"golearning-goframe/api/auth/v1"
)

type IAuthV1 interface {
	AdminLogin(ctx context.Context, req *v1.AdminLoginReq) (res *v1.AdminLoginRes, err error)
	AdminRefreshToken(ctx context.Context, req *v1.AdminRefreshTokenReq) (res *v1.AdminRefreshTokenRes, err error)
	AdminLogout(ctx context.Context, req *v1.AdminLogoutReq) (res *v1.AdminLogoutRes, err error)
}
