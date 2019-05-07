#!/usr/bin/env bash

INSTALL_DIR=/usr/local/bin
if [[ ! -e "${INSTALL_DIR}/grapi" ]]; then
    echo "Installing grapi..."
    curl -Lo grapi https://github.com/izumin5210/grapi/releases/download/v0.3.2/grapi_linux_amd64
    chmod +x grapi
    sudo mv grapi /usr/local/bin
fi