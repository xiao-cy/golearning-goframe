package model

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

const (
	// 上下文变量存储键名，前后端系统共享
	ContextKey = "ContextKey"
)

// 请求上下文结构
type Context struct {
	Session *ghttp.Session // 当前Session管理对象
}
