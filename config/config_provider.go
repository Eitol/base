package config

import (
	"base/config/env"
	"base/pkg/config_extractor"
	"github.com/joho/godotenv"
	"time"
)

// ServerConfig env names

type serverConfig struct {
	GRPCPort     int
	HTTPPort     int
	RunHTTPProxy bool
	Environment  env.Env
}

type securityConfig struct {
	PasswordEncKey   []byte
	PasswordCost     int64
	TokenSecret      []byte
	TokenDuration    time.Duration
	DisableTokenAuth bool
}

var CouchDBConfig = defaultDBConfig
var ServerCfg = defaultServerConfig
var ReporterCfg = defaultReporterConfig
var SecurityCfg = defaultSecurityConfig

// add here new configurations that need load to env
var extractorArgs = config_extractor.ExtractorArgs{
	Configs: []interface{}{
		&ReporterCfg, "REPORTER",
		&ServerCfg, "",
		&CouchDBConfig, "COUCHDB",
		&SecurityCfg, "",
	},
}

func Init(isTest bool) error {
	envPath := DefaultEnvFilePath
	if isTest {
		envPath = DefaultEnvTestFilePath
	}
	err := godotenv.Load(envPath)
	err = config_extractor.Extract(extractorArgs)
	if err != nil {
		return err
	}
	return nil
}
