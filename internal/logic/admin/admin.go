package admin

import (
	"context"
	"golearning-goframe/internal/dao"
	"golearning-goframe/internal/model"
	"golearning-goframe/internal/model/do"
	"golearning-goframe/internal/model/entity"
	"golearning-goframe/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	sAdmin struct{}
)

func init() {
	service.RegisterAdmin(New())
}

func New() service.IAdmin {
	return &sAdmin{}
}

func (s *sAdmin) CreateAdmin(ctx context.Context, AInput model.AdminCreateInput) (adminId int64, err error) {
	if AInput.AdminName == "" {
		return 0, gerror.Newf(`admin name is "", Please re-input`)
	}
	if AInput.Mobile == "" {
		return 0, gerror.Newf(`Mobile is "", Please re-input`)
	}
	if AInput.Password == "" {
		return 0, gerror.Newf(`Password  is "", Please re-input`)
	}
	if AInput.Department == "" {
		return 0, gerror.Newf(`Department  is "", Please re-select`)
	}

	is, err := s.IsMobileAvailable(ctx, AInput.Mobile)

	if err != nil {
		return 0, err
	}
	if !is {
		return 0, gerror.Newf(`管理员已注册`)
	}

	available, err := s.IsAdminNameAvailable(ctx, AInput.AdminName)
	if err != nil {
		return 0, err
	}
	if !available {
		return 0, gerror.Newf(`管理员名称已存在`)
	}

	oAdminId, err := g.Model("admin").Data(AInput).InsertAndGetId()
	if err != nil {
		return 0, err
	}

	return oAdminId, nil
}

func (s *sAdmin) GetAdminByAdminId(ctx context.Context, adminId int64) (output *entity.Admin, err error) {

	err = dao.Admin.Ctx(ctx).Where(do.Admin{
		AdminId: adminId,
	}).Scan(&output)
	g.Log().Info(ctx, output)
	if err != nil {
		return nil, err
	}
	return
}

func (s *sAdmin) GetAdminByMobileAndPassword(ctx context.Context, mobile string, password string) (admin *entity.Admin, err error) {
	err = dao.Admin.Ctx(ctx).
		Where(do.Admin{
			Mobile:   mobile,
			Password: password,
		}).
		Scan(&admin)
	g.Log().Info(ctx, admin)
	if err != nil {
		return nil, err
	}
	return
}

func (s *sAdmin) IsMobileAvailable(ctx context.Context, moblie string) (is bool, err error) {
	count, err := dao.Admin.Ctx(ctx).Where(do.Admin{
		Mobile: moblie,
	}).Count()

	if err != nil {
		return false, err
	}

	return count == 0, nil
}

func (s *sAdmin) IsAdminNameAvailable(ctx context.Context, adminName string) (is bool, err error) {
	count, err := dao.Admin.Ctx(ctx).Where(do.Admin{
		AdminName: adminName,
	}).Count()

	if err != nil {
		return false, err
	}

	return count == 0, nil
}
