package controllers

import (
	api "base/api/gen/go/api/protos"
	"base/config"
	"context"
	"github.com/gogo/protobuf/types"
	"time"

	"github.com/izumin5210/grapi/pkg/grapiserver"
)

type ServerInfoSrv struct {
	grapiserver.Server
}

func (s *ServerInfoSrv) GetInfo(ctx context.Context, empty2 *types.Empty) (*api.ServerInfo, error) {
	return &api.ServerInfo{
		Version:            config.Version,
		Compilation:        config.Compilation,
		CompilationTimeStr: config.BuildDateStr,
		ServerTimeStr:      time.Now().Format(time.RFC3339),
		ServerTimeTs:       time.Now().Unix(),
	}, nil
}
