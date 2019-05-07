#!/usr/bin/env bash
GRPC_VERSION=1.20.x
sudo apt-get install autoconf libtool clang libc++-dev build-essential pkg-config
rm -rf /tmp/grpc
git clone https://github.com/grpc/grpc.git /tmp/grpc
cd /tmp/grpc
git submodule update --init
make plugins -j12
cd /tmp/grpc/bins/opt
ls | grep grpc_ | sed 'p;s/_plugin//;s/^/protoc-gen-/' | xargs -n2 mv
sudo mv /tmp/grpc/bins/opt/protoc-gen-* /usr/local/bin/
