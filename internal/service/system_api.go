package service

import (
	"context"
	"fmt"
	"math"

	erpv1 "github.com/donech/proto-go/donech/erp/v1"

	"github.com/jinzhu/gorm"

	"github.com/donech/erp/internal/common"
	"github.com/donech/erp/internal/tool"
	"github.com/donech/tool/tabler"
	"github.com/donech/tool/xlog"

	"github.com/donech/erp/internal/domain/system"
)

type SystemAPIServer struct{}

func (s SystemAPIServer) SaveUser(ctx context.Context, request *erpv1.SaveUserRequest) (*erpv1.SaveUserResponse, error) {
	_, err := system.SaveUser(ctx, request.Account, request.Name, request.Password)
	if err != nil {
		xlog.S(ctx).Errorf("SaveUser error, request=%v, err=%v", request, err)
		return nil, err
	}
	return &erpv1.SaveUserResponse{
		Code: common.SuccessCode,
		Msg:  common.ResponseMsg(common.SuccessCode),
	}, nil
}

func (s SystemAPIServer) CheckAccount(ctx context.Context, request *erpv1.CheckAccountRequest) (*erpv1.CheckAccountResponse, error) {
	_, err := system.GetUserByAccount(ctx, request.Account)
	if gorm.IsRecordNotFoundError(err) {
		return &erpv1.CheckAccountResponse{
			Code: common.SuccessCode,
			Msg:  common.ResponseMsg(common.SuccessCode),
		}, nil
	}
	if err != nil {
		xlog.S(ctx).Errorf("system.GetUserByAccount error, request=%v, err=%v", request, err)
		return nil, err
	}
	return &erpv1.CheckAccountResponse{
		Code: common.ErrorCode,
		Msg:  "account already exist",
	}, nil
}

func (s SystemAPIServer) Users(ctx context.Context, request *erpv1.UsersRequest) (*erpv1.UsersResponse, error) {
	if request.Pager == nil {
		request.Pager = &erpv1.Pager{}
	}
	if request.Pager.PageSize == 0 {
		request.Pager.PageSize = 10
	}
	if request.Pager.PageNum == 0 && request.Pager.Cursor == 0 {
		request.Pager.Cursor = math.MaxInt64
	}
	var condition string
	if request.Name != "" {
		condition = "name like ?"
	}
	users, hasMore, err := system.Users(ctx, tabler.Pager{
		PageSize: request.Pager.PageSize,
		Cursor:   request.Pager.Cursor,
		PageNum:  request.Pager.PageNum,
	}, condition, fmt.Sprintf("%%%s%%", request.Name))
	if err != nil {
		xlog.S(ctx).Errorf("get users error, request=%v, err=%v", request, err)
		return nil, err
	}
	var data []*erpv1.UsersResponse_Data
	err = tool.JsonCopy(users, &data)
	if err != nil {
		xlog.S(ctx).Errorf("JsonCopy error, users=%v, err=%v", users, err)
		return nil, err
	}
	request.Pager.HasMore = hasMore
	return &erpv1.UsersResponse{
		Code:  common.SuccessCode,
		Msg:   common.ResponseMsg(common.SuccessCode),
		Data:  data,
		Pager: request.Pager,
	}, err
}
