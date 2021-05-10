package service

import (
	"context"

	"github.com/donech/erp/internal/common"
	"github.com/donech/erp/internal/domain/supplier"
	erpv1 "github.com/donech/proto-go/donech/erp/v1"
	"github.com/donech/tool/xlog"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type AdminAPIServer struct{}

func (s AdminAPIServer) GetSupplierList(ctx context.Context, request *erpv1.GetSupplierListRequest) (*erpv1.GetSupplierListResponse, error) {
	suppliers, err := supplier.GetSupplierByPage(ctx, nil, 1, 20)
	if err != nil {
		return nil, err
	}
	xlog.L(ctx).Debug("GetSupplierList", zap.Reflect("suppliers", suppliers))
	dto := make([]*erpv1.Supplier, 0, len(suppliers))
	err = copier.Copy(&dto, suppliers)
	if err != nil {
		return nil, err
	}
	return &erpv1.GetSupplierListResponse{
		Code:    common.SuccessCode,
		Message: common.ResponseMsg(common.SuccessCode),
		Data:    dto,
	}, nil
}
