package service

import (
	"context"
	"github.com/donech/erp/internal/common"
	"github.com/donech/erp/internal/proto"
)

type HealthCheckService struct{}

func (HealthCheckService) Readiness(ctx context.Context, req *proto.ReadinessReq) (*proto.ReadinessResp, error) {
	return &proto.ReadinessResp{
		Code: common.SuccessCode,
		Message: "pong",
	}, nil
}

func (HealthCheckService) Liveness(ctx context.Context, req *proto.LivenessReq) (*proto.LivenessResp, error) {
	return &proto.LivenessResp{
		Code: common.SuccessCode,
		Message: "pong",
	}, nil
}
