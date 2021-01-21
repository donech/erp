package service

import (
	"context"

	"github.com/donech/erp/internal/proto"
)

type GreeterService struct{}

func (GreeterService) SayHello(ctx context.Context, req *proto.HelloReq) (*proto.HelloResp, error) {
	return &proto.HelloResp{
		Message: "Hello " + req.Name,
	}, nil
}
