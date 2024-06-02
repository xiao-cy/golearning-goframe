// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"golearning-goframe/internal/model"
	"golearning-goframe/internal/model/entity"
)

type (
	IAdmin interface {
		CreateAdmin(ctx context.Context, AInput model.AdminCreateInput) (adminId int64, err error)
		GetAdminByAdminId(ctx context.Context, adminId int64) (output *entity.Admin, err error)
		IsMobileAvailable(ctx context.Context, moblie string) (is bool, err error)
		IsAdminNameAvailable(ctx context.Context, adminName string) (is bool, err error)
	}
)

var (
	localAdmin IAdmin
)

func Admin() IAdmin {
	if localAdmin == nil {
		panic("implement not found for interface IAdmin, forgot register?")
	}
	return localAdmin
}

func RegisterAdmin(i IAdmin) {
	localAdmin = i
}
