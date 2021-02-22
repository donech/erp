package internal

import (
	"github.com/donech/erp/internal/proto"
	"github.com/donech/erp/internal/service"
	"github.com/donech/tool/entry"
	"github.com/donech/tool/entry/xgrpc"
	"google.golang.org/grpc"
)

func NewGrpcEntry(cfg xgrpc.Config) entry.Entry {
	return xgrpc.New(cfg, xgrpc.WithRegisteServer(registeServer))
}

func registeServer(server *grpc.Server) {
	srv := service.GreeterService{}
	system := service.SystemService{}
	proto.RegisterGreeterServer(server, srv)
	proto.RegisterSystemServer(server, system)
}
