package models

import (
	api "base/api/gen/go/api/protos"
	"base/config"
	"base/config/constants"
	"base/config/ecodes"
	"base/pkg/cerror"
	"base/pkg/couchdb"
	"base/pkg/crypto"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
)

type User api.User

const userTypeDoc = couchdb.DDoc("users")
const (
	byEmail = couchdb.ViewName("byEmail")
)

func NewUser(user *User) cerror.Error {
	cerr := user.LoadByEmail(user.Email)
	if !cerr.IsError() {
		return ecodes.ErrDuplicatedData.AddMeta("email", user.Email)
	}
	encPass, err := crypto.EncodePassword([]byte(user.Password), config.SecurityCfg.PasswordEncKey)
	if err != nil {
		return cerror.From(err)
	}
	// to avoid corruption when in saving
	encoded := base64.StdEncoding.EncodeToString([]byte(encPass))
	user.Password = encoded
	id, _, err := couchdb.Post(user)
	if err != nil {
		return cerror.From(err)
	}
	user.Id = id
	user.cleanSensitiveData()
	return cerror.Nil
}

// TODO: Add failed attempts counter
func (u *User) Login(plainPWD string) ([]byte, cerror.Error) {
	cerr := u.LoadByEmail(u.Email)
	if cerr.IsError() {
		return nil, cerr
	}
	decPass, err := base64.StdEncoding.DecodeString(u.Password)
	isValid, err := crypto.CheckPassword([]byte(plainPWD), []byte(decPass), config.SecurityCfg.PasswordEncKey)
	if err != nil {
		return nil, cerror.From(err)
	}
	if !isValid {
		return nil, ecodes.ErrInvalidPassword
	}
	u.cleanSensitiveData()
	token, err := crypto.CreateTokenString(&crypto.Claims{
		Type: constants.User,
		StandardClaims: jwt.StandardClaims{
			Id: u.Id,
		},
	}, config.SecurityCfg.TokenSecret, config.SecurityCfg.TokenDuration)
	return []byte(token), cerror.Nil
}

func (u *User) cleanSensitiveData() {
	u.Password = ""
}

func (u *User) LoadByEmail(email string) cerror.Error {
	userResult := User{}
	cerr := couchdb.FindOneInView(userTypeDoc, byEmail, u.Email, &userResult)
	if cerr.IsError() {
		return cerr
	}
	*u = userResult
	return cerror.Nil
}
