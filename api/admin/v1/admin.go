package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type Department string

const (
	DepartmentRoot Department = "root"
	DepartmentUser Department = "user"
)

func (c Department) String() string {
	switch c {
	case DepartmentRoot:
		return "root"
	case DepartmentUser:
		return "user"
	default:
		return ""
	}
}

type GetAdminInfoReq struct {
	g.Meta  `path:"/getAdminInfo" tags:"AdminInfo" method:"get" summary:"获取用户信息"`
	AdminId int64
}
type GetAdminInfoRes struct {
	g.Meta    `mime:"text/html" example:"string"`
	AdminInfo *AdminInfo `json:"AdminInfo"`
}

type CreateAdminReq struct {
	g.Meta     `path:"/createAdminInfo" tags:"AdminInfo" method:"GET" summary:"获取用户信息"`
	AdminName  string     `v:"required|length:4,30#请输入账号|账号长度为:{min}到:{max}位"`
	Mobile     string     `v:"required|phone"`
	Password   string     `v:"required|length:6,32#请输入密码|密码长度不够"`
	Department Department `v:"required|enums"`
}

type CreateAdminRes struct {
	g.Meta  `mime:"text/html" example:"string"`
	AdminId int64
}

type AdminInfo struct {
	AdminId    int64  `json:"AdminId"`
	AdminName  string `json:"AdminName"`
	Department string `json:"Department"`
}
