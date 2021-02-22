package service

import (
	"context"

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
