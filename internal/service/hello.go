package service

import (
	"context"

	"github.com/donech/erp/internal/domain/system"
	"github.com/donech/tool/xlog"

	"github.com/donech/erp/internal/proto"
)

type GreeterService struct{}

func (GreeterService) SayHello(ctx context.Context, req *proto.HelloReq) (*proto.HelloResp, error) {
	user, err := system.AuthUser(ctx, false)
	if err != nil {
		xlog.S(ctx).Error("AuthUser error, ", err)
		return nil, err
	}
	return &proto.HelloResp{
		Message: "Hello " + req.Name + " and " + user.Name,
	}, nil
}
