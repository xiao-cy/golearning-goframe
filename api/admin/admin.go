// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package admin

import (
	"context"

	"golearning-goframe/api/admin/v1"
)

type IAdminV1 interface {
	GetAdminInfo(ctx context.Context, req *v1.GetAdminInfoReq) (res *v1.GetAdminInfoRes, err error)
	CreateAdmin(ctx context.Context, req *v1.CreateAdminReq) (res *v1.CreateAdminRes, err error)
}
