package logic

import (
	"context"
	"go-zero-admin/rpc/model/umsmodel"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type IntegrationConsumeSettingUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIntegrationConsumeSettingUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationConsumeSettingUpdateLogic {
	return &IntegrationConsumeSettingUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IntegrationConsumeSettingUpdateLogic) IntegrationConsumeSettingUpdate(in *ums.IntegrationConsumeSettingUpdateReq) (*ums.IntegrationConsumeSettingUpdateResp, error) {
	err := l.svcCtx.UmsIntegrationConsumeSettingModel.Update(umsmodel.UmsIntegrationConsumeSetting{
		Id:                 in.Id,
		DeductionPerAmount: in.DeductionPerAmount,
		MaxPercentPerOrder: in.MaxPercentPerOrder,
		UseUnit:            in.UseUnit,
		CouponStatus:       in.CouponStatus,
	})
	if err != nil {
		return nil, err
	}

	return &ums.IntegrationConsumeSettingUpdateResp{}, nil
}
