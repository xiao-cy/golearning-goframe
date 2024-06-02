package damo

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "golearning-goframe/api/damo/v1"
	"golearning-goframe/internal/dao"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func (c *ControllerV1) Damo(ctx context.Context, req *v1.DamoReq) (res *v1.DamoRes, err error) {
	// r, err := g.Client().Get(nil, "http://127.0.0.1:8000/hello")
	// if err != nil {
	// 	panic(err)
	// }
	res = &v1.DamoRes{}

	g.Log().Info(ctx, req)

	err = dao.Products.
		Ctx(ctx).
		Fields(res.ProductInfo).
		Where(dao.Products.Columns().Id, req.Id).
		Scan(&res.ProductInfo)

	// g.Model("products").Delete("id", req.Id)

	g.Log().Info(ctx, res)

	if err != nil {
		return nil, err
	}
	g.RequestFromCtx(ctx).Response.Writeln(res)
	return
}
