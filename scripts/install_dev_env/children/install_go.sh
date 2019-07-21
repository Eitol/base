#!/usr/bin/env bash

GO_VERSION=1.12
DEFAULT_GOPATH=~/go

if [[ "${GOPATH}" == "" ]]; then
    export GOPATH=${DEFAULT_GOPATH}
    echo "GOPATH=${DEFAULT_GOPATH}"
fi

if [[ $(which go) == ""  ]]; then
    sudo snap install go --classic
    mkdir ${DEFAULT_GOPATH}
    cd ${DEFAULT_GOPATH}
    mkdir bin src pkg
    cd -
fi


