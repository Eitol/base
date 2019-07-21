#!/usr/bin/env bash
GRAPI_VERSION=v0.4.0
INSTALL_DIR=/usr/local/bin

if [[ ! -e "${INSTALL_DIR}/grapi" ]]; then
    echo "Installing grapi..."
    curl -Lo grapi https://github.com/izumin5210/grapi/releases/download/${GRAPI_VERSION}/grapi_linux_amd64
    chmod +x grapi
    sudo mv grapi /usr/local/bin
fi