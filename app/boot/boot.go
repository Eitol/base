package boot

import (
	"base/config"
	"base/pkg/couchdb"
	"fmt"
	"google.golang.org/grpc/grpclog"
)

func Boot(isTest bool) {
	err := config.Init(isTest)
	if err != nil {
		grpclog.Fatal("Error loading the config: %v", err)
	}
	initDBS()
}

func initDBS() {
	_, err := couchdb.Configure(config.CouchDBConfig)
	if err != nil {
		fmt.Print(err)
	}
}
