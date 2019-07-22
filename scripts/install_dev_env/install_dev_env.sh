#!/usr/bin/env bash

GO_VERSION=1.12
DEFAULT_GOPATH=~/go

if [[ "${GOPATH}" == "" ]]; then
    export GOPATH=${DEFAULT_GOPATH}
    export GOBIN=${GOPATH}/bin
    echo "GOPATH=${DEFAULT_GOPATH}"
    export PATH=$PATH:${GOPATH}/bin
    if [[ "$(cat ~/.bashrc | grep GOPATH)" == "" ]]; then
        echo "export GOPATH=${GOPATH}" >> ~/.bashrc
        echo "export GOBIN=${GOBIN}" >> ~/.bashrc
        echo "export PATH=$PATH:${GOBIN}" >> ~/.bashrc
    fi
fi

check_if_fail(){
   if [[ $? -eq 0 ]]; then
      echo OK
   else
      echo "Fail in $1"
      exit 1
   fi
}

bash ${PWD}/scripts/install_dev_env/children/install_go.sh
check_if_fail install_go.sh
bash ${PWD}/scripts/install_dev_env/children/install_grapi.sh
check_if_fail install_grapi.sh
bash ${PWD}/scripts/install_dev_env/children/install_api_common_protos.sh
check_if_fail install_api_common_protos.sh
bash ${PWD}/scripts/install_dev_env/children/install_protoc.sh
check_if_fail install_protoc.sh
bash ${PWD}/scripts/install_dev_env/children/install_proto_plugins.sh
check_if_fail install_proto_plugins.sh
bash ${PWD}/scripts/install_dev_env/children/install_prototool.sh
check_if_fail install_prototool.sh
bash ${PWD}/scripts/install_dev_env/children/install_python_deps.sh
check_if_fail install_python_deps.sh
bash ${PWD}/scripts/install_dev_env/children/install_dart.sh
check_if_fail install_dart.sh