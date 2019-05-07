package main

import (
	"base/app/boot"
	"base/app/controllers"
	"base/app/interceptors"
	"base/config"
	"fmt"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/grpclog"
	"os"
)

func main() {
	boot.Boot(false)
	exitCode := run()
	os.Exit(exitCode)
}

// Run starts the grapiserver.
func Run() error {
	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithGrpcServerUnaryInterceptors(interceptors.AuthInterceptor),
		grapiserver.WithGrpcAddr("tcp", fmt.Sprintf("localhost:%d", config.ServerCfg.GRPCPort)),
		grapiserver.WithGatewayAddr("tcp", fmt.Sprintf("localhost:%d", config.ServerCfg.HTTPPort)),
		grapiserver.WithServers(
			&controllers.ServerInfoSrv{},
			&controllers.UserSrv{},
		),
	)
	return s.Serve()
}

func run() int {
	err := Run()
	if err != nil {
		grpclog.Errorf("server was shutdown with errors: %v", err)
		return 1
	}
	return 0
}
