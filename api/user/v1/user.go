package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetUserInfoReq struct {
	g.Meta `path:"/getUserInfo" tags:"userInfo" method:"get" summary:"获取用户信息"`
	UserId int
}
type GetUserInfoRes struct {
	g.Meta   `mime:"text/html" example:"string"`
	UserInfo *UserInfo `json:"UserInfo"`
}

type UserInfo struct {
	UserId     string `json:"UserId"`
	UserName   string `json:"UserName"`
	Department string `json:"Department"`
}
