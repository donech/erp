package service

import (
	"context"

	erpv1 "github.com/donech/proto-go/donech/erp/v1"

	"github.com/donech/tool/xlog"

	"github.com/donech/erp/internal/common"
	"github.com/donech/tool/xjwt"
)

type GateAPIServer struct{}

func (s GateAPIServer) Login(ctx context.Context, request *erpv1.LoginRequest) (*erpv1.LoginResponse, error) {
	jwtFactory := common.GetJwtFactory()
	token, err := jwtFactory.GenerateToken(ctx, xjwt.LoginForm{
		Username: request.Account,
		Password: request.Password,
	})
	if err != nil {
		xlog.S(ctx).Warn("Login error, ", err)
		return &erpv1.LoginResponse{
			Code: common.ErrorCode,
			Msg:  common.ResponseMsg(common.ErrorCode),
		}, nil
	}
	return &erpv1.LoginResponse{
		Code:  common.SuccessCode,
		Msg:   common.ResponseMsg(common.SuccessCode),
		Token: token,
	}, err
}

func (s GateAPIServer) Readiness(ctx context.Context, request *erpv1.ReadinessRequest) (*erpv1.ReadinessResponse, error) {
	return &erpv1.ReadinessResponse{
		Code:    common.SuccessCode,
		Message: "pong",
	}, nil
}

func (s GateAPIServer) Liveness(ctx context.Context, request *erpv1.LivenessRequest) (*erpv1.LivenessResponse, error) {
	return &erpv1.LivenessResponse{
		Code:    common.SuccessCode,
		Message: "pong",
	}, nil
}
