#!/usr/bin/env bash

if [[ ! -e "${GOPATH}/bin/dep" ]]; then
    echo "Installing dep..."
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
fi
