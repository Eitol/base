package config_extractor

import (
	"base/config/env"
	"github.com/joho/godotenv"
	"reflect"
	"testing"
)

const testEnv string = "test.env"

type serverConfig struct {
	GRPCPort     int  `env:"GRPC_PORT"`
	HTTPPort     int  `env:"HTTP_PORT"`
	RunHTTPProxy bool `env:"RUN_HTTP_PROXY"`
	Environment  env.Env
}

var cfg = serverConfig{
	GRPCPort:     3000,
	HTTPPort:     3001,
	RunHTTPProxy: true,
	Environment:  "DEFAULT_VALUE",
}

var cfgWant = serverConfig{
	GRPCPort:     5000,
	HTTPPort:     3001,
	RunHTTPProxy: true,
	Environment:  "DEVELOPMENT",
}

var opts = ExtractorOptions{
	TakeNameIfNoTag: true,
}

func TestExtract(t *testing.T) {
	tests := []struct {
		name    string
		envFile string
		args    ExtractorArgs
		want    []interface{}
		wantErr bool
	}{
		{
			name:    "Success",
			envFile: testEnv,
			args: ExtractorArgs{
				Options: opts,
				Configs: []interface{}{&cfg, ""},
			},
			want:    []interface{}{&cfgWant, ""},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := godotenv.Load(tt.envFile)
			if err != nil {
				t.Errorf("Extract() invalid test env file error = %v", err)
			}
			err = Extract(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Extract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i := 0; i < len(tt.args.Configs); i++ {
				if !reflect.DeepEqual(tt.args.Configs[i], tt.want[i]) {
					t.Errorf("Extract() = %v, want %v", tt.args.Configs, tt.want)
				}
			}

		})
	}
}
