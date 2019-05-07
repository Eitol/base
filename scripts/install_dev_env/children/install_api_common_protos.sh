#!/usr/bin/env bash

echo "Installing api_common_protos..."
curl -Lo api-common-protos-0.1.0.tar.gz https://codeload.github.com/googleapis/api-common-protos/tar.gz/0.1.0
tar -xzf api-common-protos-0.1.0.tar.gz
sudo cp -r api-common-protos-0.1.0/google /usr/local/include
rm -rf api-common-protos-0.1.0.tar.gz
rm -rf api-common-protos-0.1.0