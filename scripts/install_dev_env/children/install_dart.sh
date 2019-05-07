#!/usr/bin/env bash


if [[ ! -e "/etc/apt/sources.list.d/dart_stable.list" ]]; then
    #Enable HTTPS for apt.
    sudo apt-get update
    sudo apt-get install apt-transport-https

    #Get the Google Linux package signing key.
    sudo sh -c 'curl https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add -'

    #Set up the location of the stable repository.
    sudo sh -c 'curl https://storage.googleapis.com/download.dartlang.org/linux/debian/dart_stable.list > /etc/apt/sources.list.d/dart_stable.list'
    sudo apt-get update
    sudo apt-get install dart

    echo "export PATH=\$PATH:/usr/lib/dart/bin" | sudo tee -a /etc/profile
    echo "export PATH=\$PATH:/usr/lib/dart/bin" | sudo tee -a ~/.bashrc
    pub global activate protoc_plugin

    source /etc/profile
fi


if [[ ! -e "/usr/local/bin/protoc_dart_plugin/bin/protoc-gen-dart" ]]; then
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
    echo "export PATH=\$PATH:/usr/local/bin/protoc_dart_plugin/protoc_plugin/bin" | sudo tee -a /etc/profile
    echo "export PATH=\$PATH:/usr/local/bin/protoc_dart_plugin/protoc_plugin/bin" | sudo tee -a ~/.bashrc
fi

source ~/.bashrc