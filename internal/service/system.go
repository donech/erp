package service

import (
	"context"
	"fmt"
	"math"

	"github.com/jinzhu/gorm"

	"github.com/donech/erp/internal/common"
	"github.com/donech/erp/internal/tool"
	"github.com/donech/tool/tabler"
	"github.com/donech/tool/xlog"

	"github.com/donech/erp/internal/domain/system"

	"github.com/donech/erp/internal/proto"
)

type SystemService struct{}

func (SystemService) SaveUser(ctx context.Context, req *proto.SaveUserReq) (*proto.SaveUserResp, error) {
	_, err := system.SaveUser(ctx, req.Account, req.Name, req.Password)
	if err != nil {
		xlog.S(ctx).Errorf("SaveUser error, req=%v, err=%v", req, err)
		return nil, err
	}
	return &proto.SaveUserResp{
		Code: common.SuccessCode,
		Msg:  common.ResponseMsg(common.SuccessCode),
	}, nil
}

func (s SystemService) Users(ctx context.Context, req *proto.UsersReq) (*proto.UsersResp, error) {
	if req.Pager.PageSize == 0 {
		req.Pager.PageSize = 10
	}
	if req.Pager.PageSize == 0 && req.Pager.Cursor == 0 {
		req.Pager.Cursor = math.MaxInt64
	}
	var condition string
	if req.Name != "" {
		condition = "name like ?"
	}
	users, hasMore, err := system.Users(ctx, tabler.Pager{
		PageSize: req.Pager.PageSize,
		Cursor:   req.Pager.Cursor,
		PageNum:  req.Pager.PageNum,
	}, condition, fmt.Sprintf("%%%s%%", req.Name))
	if err != nil {
		xlog.S(ctx).Errorf("get users error, req=%v, err=%v", req, err)
		return nil, err
	}
	var data []*proto.UsersResp_Data
	err = tool.JsonCopy(users, &data)
	if err != nil {
		xlog.S(ctx).Errorf("JsonCopy error, users=%v, err=%v", users, err)
		return nil, err
	}
	req.Pager.HasMore = hasMore
	return &proto.UsersResp{
		Code:  common.SuccessCode,
		Msg:   common.ResponseMsg(common.SuccessCode),
		Data:  data,
		Pager: req.Pager,
	}, err
}

func (s SystemService) CheckAccount(ctx context.Context, req *proto.SaveUserReq) (*proto.SaveUserResp, error) {
	_, err := system.GetUserByAccount(ctx, req.Account)
	if gorm.IsRecordNotFoundError(err) {
		return &proto.SaveUserResp{
			Code: common.SuccessCode,
			Msg:  common.ResponseMsg(common.SuccessCode),
		}, nil
	}
	if err != nil {
		xlog.S(ctx).Errorf("system.GetUserByAccount error, req=%v, err=%v", req, err)
		return nil, err
	}
	return &proto.SaveUserResp{
		Code: common.ErrorCode,
		Msg:  "account already exist",
	}, nil
}
