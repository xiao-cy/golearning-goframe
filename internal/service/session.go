// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"golearning-goframe/internal/model/entity"
)

type (
	ISession interface {
		SetAdmin(ctx context.Context, admin *entity.Admin) error
		GetAdmin(ctx context.Context) *entity.Admin
		RemoveAdmin(ctx context.Context) error
	}
)

var (
	localSession ISession
)

func Session() ISession {
	if localSession == nil {
		panic("implement not found for interface ISession, forgot register?")
	}
	return localSession
}

func RegisterSession(i ISession) {
	localSession = i
}
