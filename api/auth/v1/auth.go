package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type AdminLoginReq struct {
	g.Meta   `path:"/admin/login" tags:"userInfo" method:"POST" summary:"登录"`
	Mobile   string
	Password string
}
type AdminLoginRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
	UserId int       `json:"UserId"`
}

type AdminRefreshTokenReq struct {
	g.Meta `path:"/admin/refresh_token" method:"POST"`
	Token  string `json:"token"`
}

type AdminRefreshTokenRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type AdminLogoutReq struct {
	g.Meta `path:"/admin/logout" method:"POST"`
	Token  string `json:"token"`
}

type AdminLogoutRes struct {
}
