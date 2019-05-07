package controllers

import (
	api "base/api/gen/go/api/protos"
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// RegisterWithServer implements grapiserver.Server.RegisterWithServer.
func (s *ServerInfoSrv) RegisterWithServer(grpcSvr *grpc.Server) {
	api.RegisterServerInfoServiceServer(grpcSvr, s)
}

// RegisterWithHandler implements grapiserver.Server.RegisterWithHandler.
func (s *ServerInfoSrv) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return api.RegisterServerInfoServiceHandler(ctx, mux, conn)
}

// RegisterWithServer implements grapiserver.Server.RegisterWithServer.
func (s *UserSrv) RegisterWithServer(grpcSvr *grpc.Server) {
	api.RegisterUserServiceServer(grpcSvr, s)
}

// RegisterWithHandler implements grapiserver.Server.RegisterWithHandler.
func (s *UserSrv) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return api.RegisterUserServiceHandler(ctx, mux, conn)
}
