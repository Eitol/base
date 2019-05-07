package main

import (
	"base/api/gen/go/api/protos"
	"context"
	"fmt"
	"github.com/gogo/protobuf/types"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:3001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := genpb.NewServerInfoServiceClient(conn)
	data, err := client.GetInfo(context.Background(), &types.Empty{})

	fmt.Println(data)

}
