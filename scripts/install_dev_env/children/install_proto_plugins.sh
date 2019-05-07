#!/usr/bin/env bash
echo "Installing proto plugins..."
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
go get github.com/gogo/protobuf/protoc-gen-gogofast
go get github.com/gogo/protobuf/protoc-gen-gogofaster
go get github.com/gogo/protobuf/protoc-gen-gogoslick

# Python support
if [[ ! -e "python3" ]]; then
    sudo apt install python3 python3-pip
fi
sudo pip3 install grpcio
sudo pip3 install grpcio-tools
sudo pip3 install googleapis-common-protos
