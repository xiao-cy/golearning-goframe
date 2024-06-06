package bizctx

import (
	"context"

	"golearning-goframe/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"

	"golearning-goframe/internal/model"
)

type sBizCtx struct{}

func init() {
	service.RegisterBizCtx(New())
}

func New() service.IBizCtx {
	return &sBizCtx{}
}

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *sBizCtx) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(model.ContextKey, customCtx)
}

// Get 获得上下文变量，如果没有设置，那么返回nil
func (s *sBizCtx) Get(ctx context.Context) *model.Context {
	value := ctx.Value(model.ContextKey)
	if value == nil {
		return nil
	}

	// g.Log().Line(true).Info(ctx, value)
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}
