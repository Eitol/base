package models

import (
	"base/app/boot"
	"base/pkg/cerror"
	"reflect"
	"testing"
)

func init() {
	boot.Boot(true)
}

var user = User{
	Username: "eitol",
	Email:    "hector.oliveros.leon@gmail.com",
	Password: "123456",
}

func TestNewUser(t *testing.T) {
	type args struct {
		user User
	}
	tests := []struct {
		name    string
		args    args
		wantErr cerror.Error
	}{
		{
			name: "new",
			args: args{
				user: user,
			},
			wantErr: cerror.Error{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := NewUser(&tt.args.user)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("NewUser() got = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestUser_Login(t *testing.T) {
	type args struct {
		user *User
	}
	tests := []struct {
		name    string
		args    args
		wantErr cerror.Error
	}{
		{
			name: "new",
			args: args{
				user: &user,
			},
			wantErr: cerror.Error{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, got := tt.args.user.Login(tt.args.user.Password)
			if !reflect.DeepEqual(got, tt.wantErr) {
				t.Errorf("User.Login() = %v, want %v", got, tt.wantErr)
			}
			if len(token) == 0 {
				t.Errorf("User.Login() = invalid token len")
			}
		})
	}
}
