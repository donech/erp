package internal

import (
	"context"
	"errors"

	"github.com/donech/erp/internal/common"

	"github.com/donech/erp/internal/domain/system"

	"github.com/dgrijalva/jwt-go"

	"github.com/donech/erp/internal/proto"
	"github.com/donech/erp/internal/service"
	"github.com/donech/tool/entry"
	"github.com/donech/tool/entry/xgrpc"
	"github.com/donech/tool/xjwt"
	"google.golang.org/grpc"
)

func NewGrpcEntry(cfg xgrpc.Config) entry.Entry {
	common.InitJwtFactory(Login)
	return xgrpc.New(cfg, xgrpc.WithRegisteServer(registeServer), xgrpc.WithJwtFactory(common.GetJwtFactory()))
}

func registeServer(server *grpc.Server) {
	srv := service.GreeterService{}
	systemSrv := service.SystemService{}
	commonSrv := service.CommonService{}
	proto.RegisterGreeterServer(server, srv)
	proto.RegisterSystemServer(server, systemSrv)
	proto.RegisterCommonServer(server, commonSrv)
}

func Login(ctx context.Context, form xjwt.LoginForm) (jwt.MapClaims, error) {
	user, err := system.GetUserByAccount(ctx, form.Username)
	if err != nil {
		return nil, err
	}

	if !common.ValidatePassword(form.Password, user.Password) {
		return nil, errors.New("username or password error")
	}

	return jwt.MapClaims{
		"id":   user.ID,
		"name": user.Name,
	}, nil
}
