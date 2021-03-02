package service

import (
	"context"

	"github.com/donech/tool/xlog"

	"github.com/donech/erp/internal/common"
	"github.com/donech/tool/xjwt"

	"github.com/donech/erp/internal/proto"
)

type CommonService struct{}

func (CommonService) Login(ctx context.Context, req *proto.LoginReq) (*proto.LoginResp, error) {
	jwtFactory := common.GetJwtFactory()
	token, err := jwtFactory.GenerateToken(ctx, xjwt.LoginForm{
		Username: req.Account,
		Password: req.Password,
	})
	if err != nil {
		xlog.S(ctx).Warn("Login error, ", err)
		return &proto.LoginResp{
			Code: common.ErrorCode,
			Msg:  common.ResponseMsg(common.ErrorCode),
		}, nil
	}
	return &proto.LoginResp{
		Code:  common.SuccessCode,
		Msg:   common.ResponseMsg(common.SuccessCode),
		Token: token,
	}, err
}
