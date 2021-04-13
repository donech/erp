package service

import (
	"context"
	"fmt"
	"math"

	"github.com/donech/erp/internal/tool"
	"github.com/donech/tool/tabler"

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
		return nil, err
	}
	var data []*proto.UsersResp_Data
	err = tool.JsonCopy(users, &data)
	if err != nil {
		return nil, err
	}
	req.Pager.HasMore = hasMore
	return &proto.UsersResp{
		Code:  0,
		Msg:   "success",
		Data:  data,
		Pager: req.Pager,
	}, err
}
