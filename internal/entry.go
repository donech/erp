package internal

import (
	"context"
	"errors"
	"strconv"

	erpv1 "github.com/donech/proto-go/donech/erp/v1"

	"github.com/donech/erp/internal/common"

	"github.com/donech/erp/internal/domain/system"

	"github.com/dgrijalva/jwt-go"

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
		"/donech.erp.v1.GateAPI/Readiness": true,
		"/donech.erp.v1.GateAPI/Liveness":  true,
		"/donech.erp.v1.GateAPI/Login":     true,
	}
}

func registeServer(server *grpc.Server) {
	systemSrv := service.SystemAPIServer{}
	gateSrv := service.GateAPIServer{}
	admSrv := service.AdminAPIServer{}
	erpv1.RegisterSystemAPIServer(server, systemSrv)
	erpv1.RegisterAdminAPIServer(server, admSrv)
	erpv1.RegisterGateAPIServer(server, gateSrv)

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
