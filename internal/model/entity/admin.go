// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Admin is the golang structure for table admin.
type Admin struct {
	AdminId    uint64      `json:"adminId"    orm:"adminId"    ` //
	Mobile     string      `json:"mobile"     orm:"mobile"     ` //
	Password   string      `json:"password"   orm:"password"   ` //
	Department string      `json:"department" orm:"department" ` //
	AdminName  string      `json:"adminName"  orm:"adminName"  ` //
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at" ` //
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at" ` //
}
