package session

import (
	"context"
	"golearning-goframe/internal/model/entity"
	"golearning-goframe/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type sSession struct{}

const (
	adminSessionKey = "adminSessionKye"
)

func init() {
	service.RegisterSession(New())
}

func New() service.ISession {
	return &sSession{}
}

func (s *sSession) SetAdmin(ctx context.Context, admin *entity.Admin) error {
	g.Log().Line(true).Print(ctx, "SetAdmin")
	return service.BizCtx().Get(ctx).Session.Set(adminSessionKey, admin)
}

func (s *sSession) GetAdmin(ctx context.Context) *entity.Admin {
	g.Log().Line(true).Print(ctx, "GetAdmin")
	sessionAdmin, _ := service.BizCtx().Get(ctx).Session.Get(adminSessionKey)
	if !sessionAdmin.IsNil() {
		var admin *entity.Admin
		sessionAdmin.Struct(&admin)
		return admin
	}
	return &entity.Admin{}
}

func (s *sSession) RemoveAdmin(ctx context.Context) error {
	return service.BizCtx().Get(ctx).Session.Remove(adminSessionKey)
}
