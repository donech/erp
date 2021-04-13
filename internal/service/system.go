package service

import (
	"context"
	"github.com/donech/erp/internal/tool"
	"github.com/donech/tool/tabler"
	"math"

	"github.com/donech/erp/internal/domain/system"

	"github.com/donech/erp/internal/proto"
)

type SystemService struct{}

func (SystemService) CreateUser(ctx context.Context, req *proto.CreateUserReq) (*proto.CreateUserResp, error) {
	_, err := system.CreateUser(ctx, req.Account, req.Name, req.Password)
	if err != nil {
		return nil, err
	}
	return &proto.CreateUserResp{
		Code: 0,
		Msg:  "创建用户成功",
	}, nil
}

func (s SystemService) Users(ctx context.Context, req *proto.UsersReq) (*proto.UsersResp, error) {
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	if req.Cursor == 0 {
		req.Cursor = math.MaxInt64
	}
	users, err := system.Users(ctx, tabler.Pager{
		PageSize: req.PageSize,
		Cursor:   req.Cursor,
		PageNum: req.PageNum,
	})
	if err != nil {
		return nil, err
	}
	var data []*proto.UsersResp_Data
	err = tool.JsonCopy(users, &data)
	if err != nil {
		return nil, err
	}
	return &proto.UsersResp{
		Code:                 0,
		Msg:                  "success",
		Data:                 data,
	}, err
}

