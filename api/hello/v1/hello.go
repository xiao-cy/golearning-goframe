package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HelloReq struct {
	g.Meta `path:"/hello" tags:"Hello" method:"get" summary:"hello World"`
}
type HelloRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
