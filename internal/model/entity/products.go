// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Products is the golang structure for table products.
type Products struct {
	Id        uint64      `json:"id"        orm:"id"         ` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` //
	Code      string      `json:"code"      orm:"code"       ` //
	Price     uint64      `json:"price"     orm:"price"      ` //
}
