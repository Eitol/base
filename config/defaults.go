package config

import (
	"base/config/env"
	"base/pkg/couchdb"
	"base/pkg/reporter"
	"math"
)

const DefaultEnvFilePath = "config/env/.env"
const DefaultEnvTestFilePath = "config/env/.env_test"

var defaultServerConfig = serverConfig{
	GRPCPort:     3000,
	HTTPPort:     3001,
	RunHTTPProxy: true,
	Environment:  env.Development,
}

var defaultDBConfig = couchdb.ClientConfig{
	Host:               "http://127.0.0.1:5984/",
	User:               "admina",
	Pass:               "123456",
	PintOnCreate:       true,
	DatabaseName:       ProjectName,
	CreateDBIfNotExist: true,
}

var defaultReporterConfig = reporter.Config{
	DSN:          "",
	ProjectName:  ProjectName,
	Compilation:  Compilation,
	Version:      Version,
	BuildDateStr: BuildDateStr,
	Environment:  string(defaultServerConfig.Environment),
}

var defaultSecurityConfig = securityConfig{
	PasswordEncKey:   []byte("CHANGE_ME_ABCDEFGHIJKLMNOPQRSTUV"),
	PasswordCost:     0,
	TokenSecret:      []byte("CHANGE_ME"),
	TokenDuration:    math.MaxInt64,
	DisableTokenAuth: false,
}
