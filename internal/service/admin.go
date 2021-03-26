package service

import (
	"context"
	"github.com/donech/erp/internal/common"
	"github.com/donech/erp/internal/domain/supplier"
	"github.com/donech/erp/internal/proto"
	"github.com/donech/tool/xlog"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type AdminService struct{}

func (AdminService) GetSupplierList(ctx context.Context, req *proto.SupplierListReq) (*proto.SupplierListResp, error) {
	suppliers, err := supplier.GetSupplierByPage(ctx, nil, 1, 20);
	if err != nil {
		return nil, err
	}
	xlog.L(ctx).Debug("GetSupplierList", zap.Reflect("suppliers", suppliers))
	dto := make([]*proto.Supplier, 0, len(suppliers))
	err = copier.Copy(&dto, suppliers)
	if err != nil {
		return nil, err
	}
	return &proto.SupplierListResp{
		Code:                 common.SuccessCode,
		Message:              common.ResponseMsg(common.SuccessCode),
		Data:                 dto,
	}, nil
}
