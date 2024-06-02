package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DamoReq struct {
	g.Meta `path:"/damo" tags:"Damo" method:"get" summary:"一个练习damo"`
	Id     int `v:required`
}

type DamoRes struct {
	g.Meta      `mime:"text/html" example:"string"`
	ProductInfo *ProductInfo `protobuf:"bytes,1,opt,name=ProductInfo,proto1" json:"ProductInfo,omitempty"`
}

type ProductInfo struct {
	Id    string `protobuf:"bytes,1,opt,name=Id,proto1" json:"Id,omitempty"`
	Code  string `protobuf:"bytes,2,opt,name=Code,proto2" json:"Code,omitempty"`
	Price string `protobuf:"bytes,3,opt,name=Price,proto3" json:"Price,omitempty"`
}
