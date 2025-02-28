package department

import (
	"PowerX/internal/svc"
	"PowerX/internal/types"
	"PowerX/internal/uc/powerx"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDepartmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDepartmentLogic {
	return &CreateDepartmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDepartmentLogic) CreateDepartment(req *types.CreateDepartmentRequest) (resp *types.CreateDepartmentReply, err error) {
	dep := powerx.Department{
		Name:        req.DepName,
		PId:         req.PId,
		LeaderId:    req.LeaderId,
		Desc:        req.Desc,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Remark:      req.Remark,
	}

	err = l.svcCtx.PowerX.Organization.CreateDepartment(l.ctx, &dep)
	if err != nil {
		return nil, err
	}

	return &types.CreateDepartmentReply{
		Id: dep.ID,
	}, nil
}
