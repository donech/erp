package internal

import (
	"context"
	"errors"
	"strconv"

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
	return xgrpc.New(cfg,
		xgrpc.WithRegisteServer(registeServer),
		xgrpc.WithJwtFactory(common.GetJwtFactory()),
		xgrpc.WithJumpMethods(GetJumpMethods()),
	)
}

//GetJumpMethods 不进行 jwt 验证的 grpc handle
func GetJumpMethods() map[string]bool {
	return map[string]bool{
		"/proto.HealthCheck/Liveness":true,
		"/proto.HealthCheck/Readiness":true,
		"/proto.Common/Login":true,
	}
}

func registeServer(server *grpc.Server) {
	srv := service.HealthCheckService{}
	systemSrv := service.SystemService{}
	commonSrv := service.CommonService{}
	admSrv := service.AdminService{}
	proto.RegisterHealthCheckServer(server, srv)
	proto.RegisterSystemServer(server, systemSrv)
	proto.RegisterCommonServer(server, commonSrv)
	proto.RegisterAdminServer(server, admSrv)
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
		"id":   strconv.FormatInt(user.ID, 10),
		"name": user.Name,
	}, nil
}
