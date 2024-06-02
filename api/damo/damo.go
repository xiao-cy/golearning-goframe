// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package damo

import (
	"context"

	"golearning-goframe/api/damo/v1"
)

type IDamoV1 interface {
	Damo(ctx context.Context, req *v1.DamoReq) (res *v1.DamoRes, err error)
}
