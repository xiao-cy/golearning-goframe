// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Admin is the golang structure of table admin for DAO operations like Where/Data.
type Admin struct {
	g.Meta     `orm:"table:admin, do:true"`
	AdminId    interface{} //
	Mobile     interface{} //
	Password   interface{} //
	Department interface{} //
	AdminName  interface{} //
	UpdatedAt  *gtime.Time //
	DeletedAt  *gtime.Time //
}
