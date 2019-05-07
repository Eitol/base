package crypto

import (
	"base/config/ecodes"
	"base/pkg/cerror"
	"fmt"
	"reflect"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

const (
	validAesKey   = "AES256Key-32Characters1234567890"
	validPassword = "kamisama"
)

var bcryptAndAesCodec = map[string][]byte{
	validPassword: {
		74, 53, 66, 103, 16, 108, 85, 60, 247, 242, 65, 148,
		188, 182, 130, 18, 27, 134, 8, 99, 229, 16, 116, 190,
		223, 66, 205, 50, 29, 8, 34, 144, 37, 133, 56, 54, 56,
		246, 115, 232, 196, 197, 219, 64, 213, 147, 142, 159,
		208, 163, 207, 34, 164, 182, 6, 151, 135, 29, 247, 17,
		163, 226, 22, 15, 109, 68, 16, 149, 8, 12, 107, 154,
		196, 198, 252, 183, 43, 138, 28, 28, 197, 40, 9, 25,
		6, 182, 139, 152,
	},
}

var bcryptCodec = map[string][]byte{
	validPassword: []byte("$2a$04$OIskz3UiTF5rJ8VRupoWkODObPAHQu/R9ruviaNHIZDZnar1piPRm"),
}

var aesCodec = map[string][]byte{
	validPassword: {
		236, 15, 157, 121, 19, 242, 119, 163, 171, 215, 252,
		109, 85, 10, 217, 2, 32, 72, 20, 9, 2, 175, 199, 20,
		175, 241, 75, 29, 62, 81, 65, 112, 139, 188, 21, 103,
	},
}

func TestSetHashCost(t *testing.T) {
	type args struct {
		f func() int
	}
	tests := []struct {
		name     string
		args     args
		expected int
	}{
		{
			name: "change",
			args: args{
				f: func() int {
					return 8
				},
			},
			expected: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetHashCost(tt.args.f)
			if tt.expected != HashCost() {
				t.Errorf("SetHashCost(%d) fail", tt.expected)
			}
		})
	}
}

func TestEncodePassword(t *testing.T) {
	type args struct {
		plainPWD []byte
		key      []byte
	}
	tests := []struct {
		name               string
		args               args
		wantErrInEnc       bool
		wantErrInDec       bool
		decodingAlteration []byte
	}{
		{
			name: "ok",
			args: args{
				plainPWD: []byte(validPassword),
				key:      []byte(validAesKey),
			},
			wantErrInEnc: false,
			wantErrInDec: false,
		},
		{
			name: "bad key len",
			args: args{
				plainPWD: []byte(validPassword),
				key:      []byte("BAD KEY"),
			},
			decodingAlteration: []byte(""),
			wantErrInEnc:       true,
			wantErrInDec:       true,
		},
		{
			name: "mismatch",
			args: args{
				plainPWD: []byte(validPassword),
				key:      []byte(validAesKey),
			},
			decodingAlteration: []byte("ALTERED STRING"),
			wantErrInEnc:       false,
			wantErrInDec:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodePassword(tt.args.plainPWD, tt.args.key)
			if (err != nil) != tt.wantErrInEnc {
				t.Errorf("EncodePassword() error = %v, wantErr %v", err, tt.wantErrInEnc)
				return
			}
			fmt.Print(string(got))
			if len(tt.decodingAlteration) > 0 {
				got = append(got, tt.decodingAlteration...)
			}
			isValid, err := CheckPassword(tt.args.plainPWD, got, tt.args.key)
			if (err != nil) != tt.wantErrInDec || (isValid) == tt.wantErrInDec {
				t.Errorf("EncodePassword() decoding -> = %v", got)
			}
		})
	}
}

func TestCheckPassword(t *testing.T) {
	type args struct {
		plainPWD []byte
		encPWD   []byte
		key      []byte
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				plainPWD: []byte(validPassword),
				encPWD:   bcryptAndAesCodec[validPassword],
				key:      []byte(validAesKey),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "bad input passwd",
			args: args{
				plainPWD: []byte(validPassword + "QWERTY123"),
				encPWD:   bcryptAndAesCodec[validPassword],
				key:      []byte(validAesKey),
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckPassword(tt.args.plainPWD, tt.args.encPWD, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hashAndSalt(t *testing.T) {
	type args struct {
		pwd []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				pwd: []byte(validPassword),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hashAndSalt(tt.args.pwd)
			err := bcrypt.CompareHashAndPassword(got, []byte(validPassword))
			if err != nil {
				t.Errorf("hashAndSalt()")
			}
		})
	}
}

func Test_encrypt_decrypt(t *testing.T) {
	type args struct {
		plaintext []byte
		key       []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "",
			args: args{
				plaintext: []byte(validPassword),
				key:       []byte(validAesKey),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encrypt(tt.args.plaintext, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			dec, err := decrypt(got, tt.args.key)
			if err != nil || !reflect.DeepEqual(tt.args.plaintext, dec) {
				t.Errorf("encrypt() -> decript error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestValidateToken(t *testing.T) {
	type args struct {
		tokenString []byte
		secret      []byte
	}
	tests := []struct {
		name string
		args args
		want cerror.Error
	}{
		{
			name: "valid",
			args: args{
				// expires at Thursday, 17 August 2311 19:36:55
				tokenString: []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidXNlciIsImV4cCI6MTA3ODA2MzA2MTUsImlhdCI6MTU1NzI1ODU3OX0.1mkliOzfrNna2Cm4NGdDsh8m-PTavEbVRm9euxHbXTI"),
				secret:      []byte("CHANGE_ME"),
			},
			want: ecodes.Nil,
		},
		{
			name: "expired",
			args: args{
				tokenString: []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidXNlciIsImV4cCI6MTU1NzI1MTk1MywiaWF0IjoxNTU3MjUxOTUzfQ.cFok0-kwGlgyKqoAucP5mcH_Muz090C7eXQ2Fz992v8"),
				secret:      []byte("CHANGE_ME"),
			},
			want: ecodes.ErrTokenExpired,
		},
		{
			name: "invalid secret",
			args: args{
				// expires at Thursday, 17 August 2311 19:36:55
				tokenString: []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidXNlciIsImV4cCI6MTA3ODA2MzA2MTUsImlhdCI6MTU1NzI1ODU3OX0.1mkliOzfrNna2Cm4NGdDsh8m-PTavEbVRm9euxHbXTI"),
				secret:      []byte("CHANGE_ME_HE_HE_HE_HE"),
			},
			want: ecodes.ErrInvTokenSignature,
		},
		{
			name: "token malformed",
			args: args{
				// expires at Thursday, 17 August 2311 19:36:55
				tokenString: []byte("Im the bad token"),
				secret:      []byte("CHANGE_ME"),
			},
			want: ecodes.ErrTokenMalformed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateToken(tt.args.tokenString, tt.args.secret)
			if got.InternalCode != tt.want.InternalCode {
				t.Errorf("ValidateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
