package service

import (
	"context"
	"fmt"
	erpv1 "github.com/donech/proto-go/donech/erp/v1"
	"math"

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

func (s SystemAPIServer) CurrentUser(ctx context.Context, request *erpv1.CurrentUserRequest) (*erpv1.CurrentUserResponse, error) {
	user, err := system.AuthUser(ctx, true)
	if err != nil {
		xlog.S(ctx).Errorf("get auth user err=%v", err)
	}
	return &erpv1.CurrentUserResponse{
		Code:  common.SuccessCode,
		Msg:   common.ResponseMsg(common.SuccessCode),
		Data:  &erpv1.CurrentUserResponse_Data{
			Id:                   user.ID,
			Name:                 user.Name,
			Avatar:               "https://www.icode9.com/i/l/?n=18&i=blog/372674/201909/372674-20190925093845264-1858102286.png",
			Email:                "solarpwx@yeah.net",
			Title:                "CTO",
			Phone:                "18001023261",
			CreatedTime:          user.CreatedTime.Format("2006-01-02 13:04:05"),
		},
	}, err
}
