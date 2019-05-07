package crypto

// Hector Oliveros - 2019
// hector.oliveros.leon@gmail.com

import (
	"base/config/ecodes"
	"base/pkg/cerror"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"io"
	"time"
)

const DefaultHashCost = bcrypt.MinCost
const JWTSigningMethod = "HS256"

var HashCost = func() int { return DefaultHashCost }

type Claims struct {
	Type string `json:"name"`
	jwt.StandardClaims
}

func ValidateToken(tokenOrig []byte, secret []byte) cerror.Error {
	token, err := jwt.Parse(string(tokenOrig), func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err == nil && token.Valid {
		return cerror.Nil
	}
	ve, ok := err.(*jwt.ValidationError)
	if !ok || ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
		return ecodes.ErrInvTokenSignature.AddMeta("token", string(tokenOrig))
	}
	if ve.Errors&jwt.ValidationErrorMalformed != 0 {
		return ecodes.ErrTokenMalformed.AddMeta("token", string(tokenOrig))
	}

	if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
		// Token is either expired or not active yet
		return ecodes.ErrTokenExpired
	}
	return ecodes.ErrTokenMalformed.SetReason("can not handle").AddMeta("token", string(tokenOrig))
}

func CreateTokenString(c *Claims, secret []byte, d time.Duration) (string, error) {
	c.ExpiresAt = time.Now().Add(d).Unix()
	c.IssuedAt = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.GetSigningMethod(JWTSigningMethod), c)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func SetHashCost(f func() int) {
	if f == nil {
		return
	}
	HashCost = f
}

// hash and encode password
func EncodePassword(plainPWD []byte, key []byte) ([]byte, error) {
	hash := hashAndSalt(plainPWD)
	encHash, err := encrypt(hash, key)
	if err != nil {
		return encHash, err
	}
	return encHash, nil
}

// Compare one plan text password with encrypted password
// if encPWD is not the hash of the plainPWD then return false and nil error
// If there is an error in the process, then the error returns.
// if this happens it is possible that it is an attack on the system and an alert should occur
func CheckPassword(plainPWD []byte, encPWD []byte, key []byte) (bool, error) {
	decPWD, err := decrypt(encPWD, key)
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword(decPWD, plainPWD)
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, err
}

func hashAndSalt(pwd []byte) []byte {
	// Use GenerateFromPassword to hash & salt pwd
	hash, err := bcrypt.GenerateFromPassword(pwd, HashCost())
	if err != nil {
		// TODO: Sentry
		return []byte{}
	}
	return hash
}

func encrypt(plaintext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short. Must be at least %d", nonceSize)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
