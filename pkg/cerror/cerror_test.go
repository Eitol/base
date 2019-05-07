package cerror

import (
	"reflect"
	"testing"

	"google.golang.org/grpc/codes"
)

var (
	errLevelX = Error{
		Type:         codes.InvalidArgument,
		InternalCode: "OTHER_WEIRD_ERROR",
		Text:         "Unknown Error",
		Description:  "Test weird error",
		Reason:       "Test reason",
		ComesFrom:    nil,
	}

	errLevel0 = Error{
		Type:         codes.InvalidArgument,
		InternalCode: "EIA",
		Text:         "Invalid argument",
		Description:  "The argument 'password' is invalid",
		Reason:       "'{1}' lower than {2}",
		Meta: map[string]string{
			"arg": "'password",
			"req": "len > 6",
		},
		ComesFrom: nil,
	}

	errLevel1 = Error{
		Type:         codes.InvalidArgument,
		InternalCode: "ECPw",
		Text:         "Error checking password",
		Description:  "The argument 'password' is invalid",
		Reason:       "'password' is invalid",
		Meta:         nil,
		ComesFrom:    &errLevel0,
	}

	errLevel2 = Error{
		Type:         codes.PermissionDenied,
		InternalCode: "LOGIN_ERR",
		Text:         "login error",
		Description:  "The login process is invalid",
		Reason:       "",
		Meta: map[string]string{
			"sample_meta": "test",
		},
		ComesFrom: &errLevel1,
	}
)

func TestError_GetParents(t *testing.T) {
	tests := []struct {
		name string
		err  Error
		want []*Error
	}{
		{
			name: "with parent",
			err:  errLevel2,
			want: []*Error{&errLevel1, errLevel1.ComesFrom},
		},
		{
			name: "with no parent",
			err:  errLevel0,
			want: []*Error{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.err.GetParents(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Error.GetParents() = %v, wantComesFrom %v", got, tt.want)
			}
		})
	}
}

func TestError_IsError(t *testing.T) {
	tests := []struct {
		name string
		err  Error
		want bool
	}{
		{
			name: "is error",
			err:  errLevel0,
			want: true,
		},
		{
			name: "not error",
			err:  Error{},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.IsError(); got != tt.want {
				t.Errorf("Error.IsError() = %v, wantComesFrom %v", got, tt.want)
			}
		})
	}
}

func TestError_Equals(t *testing.T) {
	tests := []struct {
		name string
		err1 Error
		err2 Error
		want bool
	}{
		{
			name: "is error",
			err1: errLevel0,
			err2: errLevel0,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err1.Equals(tt.err2); got != tt.want {
				t.Errorf("Error.IsError() = %v, wantComesFrom %v", got, tt.want)
			}
		})
	}
}

func TestError_SetParam(t *testing.T) {
	tests := []struct {
		name   string
		err    Error
		params map[string]int
		want   string
	}{
		{
			name:   "with params",
			err:    errLevel0,
			params: map[string]int{"password": 1, "6": 2},
			want:   "'password' lower than 6",
		},
		{
			name:   "wrong param",
			err:    errLevel0,
			params: map[string]int{"password": 7, "6": 2},
			want:   "'{1}' lower than 6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for val, pos := range tt.params {
				tt.err = tt.err.SetParam(pos, val)
			}
			if tt.err.Reason != tt.want {
				t.Errorf("Error.SetParam() = %v, wantComesFrom %v", tt.err.Reason, tt.want)
			}
		})
	}
}

func TestError_From(t *testing.T) {
	tests := []struct {
		name          string
		err1          Error
		err2          Error
		wantComesFrom Error
	}{
		{
			name:          "is error",
			err1:          errLevel0,
			err2:          errLevelX,
			wantComesFrom: errLevelX,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err1.From(tt.err2); !reflect.DeepEqual(*got.ComesFrom, tt.wantComesFrom) {
				t.Errorf("Error.From() = %v, wantComesFrom %v", got, tt.wantComesFrom)
			}
			if tt.err1.ComesFrom != nil {
				t.Errorf("Error.From() = %s, original comesFrom is not nil, %v", tt.name, tt.wantComesFrom)
			}
		})
	}
}
