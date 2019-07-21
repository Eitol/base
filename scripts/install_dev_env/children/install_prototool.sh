#!/usr/bin/env bash


echo "Installing prototool..."
set -euo pipefail

TMP="$(mktemp -d)"
trap 'rm -rf "${TMP}"' EXIT
cd "${TMP}"

echo "GOPATH=${GOPATH}"

GO111MODULE=on GOBIN=${GOPATH}/bin go get github.com/uber/prototool/cmd/prototool@dev
sudo chmod +x ${GOPATH}/bin/prototool