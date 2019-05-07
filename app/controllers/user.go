package controllers

import (
	"base/api/gen/go/api/protos"
	"base/app/models"
	"context"
	"github.com/izumin5210/grapi/pkg/grapiserver"
)

type UserSrv struct {
	grapiserver.Server
}

func (u *UserSrv) Login(ctx context.Context, req *genpb.LoginReq) (*genpb.LoginResp, error) {
	user := &models.User{
		Email: req.Email,
	}
	token, cerr := user.Login(req.Password)
	if cerr.IsError() {
		return nil, cerr
	}
	return &genpb.LoginResp{
		User:  (*genpb.User)(user),
		Token: string(token),
	}, nil
}

func (u *UserSrv) SignUp(ctx context.Context, req *genpb.SignUpReq) (*genpb.SignUpResp, error) {
	cerr := models.NewUser((*models.User)(req.User))
	if cerr.IsError() {
		return nil, cerr
	}
	return &genpb.SignUpResp{
		User: req.User,
	}, nil
}
