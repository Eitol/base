package env

import "strings"

type Env string

const (
	Test        = Env("test")
	Development = Env("development")
	Staging     = Env("staging")
	Production  = Env("production")
)

var ValidEnv = []Env{Development, Staging, Production}

func IsValidEnv(envStr string) bool {
	envStr = strings.ToLower(envStr)
	for _, _env := range ValidEnv {
		if string(_env) == envStr {
			return true
		}
	}
	return false
}
