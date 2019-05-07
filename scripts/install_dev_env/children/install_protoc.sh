#!/usr/bin/env bash

echo "Installing protoc..."
if [[ ! -f /usr/local/bin/protoc ]]; then
    PROTOC_VERSION=3.7.1
    PROTOC_ZIP=protoc-${PROTOC_VERSION}-linux-x86_64.zip
    curl -OL https://github.com/google/protobuf/releases/download/v${PROTOC_VERSION}/$PROTOC_ZIP
    sudo unzip -o $PROTOC_ZIP -d /usr/local/bin/protoc
    rm -f $PROTOC_ZIP
fi