#!/usr/bin/env bash

DARTBIN=/usr/lib/dart/bin
DART_PLUG_BIN=/usr/local/bin/protoc_dart_plugin/protoc_plugin/bin
DART_PUB_HOME=~/.pub-cache/bin

if [[ ! -e "/etc/apt/sources.list.d/dart_stable.list" ]]; then
    #Enable HTTPS for apt.
    sudo apt-get update
    sudo apt-get install apt-transport-https -y

    #Get the Google Linux package signing key.
    sudo sh -c 'curl https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add -'

    #Set up the location of the stable repository.
    sudo sh -c 'curl https://storage.googleapis.com/download.dartlang.org/linux/debian/dart_stable.list > /etc/apt/sources.list.d/dart_stable.list'
    sudo apt-get update
    sudo apt-get install dart -y
    pub global activate protoc_plugin
    .pub-cache/bin
    source /etc/profile
fi


if [[ ! -e "${DART_PLUG_BIN}" ]]; then
    sudo rm -f /usr/local/bin/protoc-gen-dart
    mkdir /tmp/dart
    cd /tmp/dart
    if [[ ! -e "./protoc_dart_plugin" ]]; then
        git clone https://github.com/dart-lang/protobuf.git protoc_dart_plugin
    fi
    cd protoc_dart_plugin/protoc_plugin
    pub install
    cd ../../
    sudo cp -r protoc_dart_plugin /usr/local/bin/

fi

if [[ "$(cat ~/.bashrc | grep ${DARTBIN})" == "" ]]; then
        echo "export PATH=$PATH:${DARTBIN}:${DART_PLUG_BIN};${DART_PUB_HOME}" >> ~/.bashrc
fi

source ~/.bashrc