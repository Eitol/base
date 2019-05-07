#!/usr/bin/env bash

bash ${PWD}/scripts/install_dev_env/children/install_grapi.sh
bash ${PWD}/scripts/install_dev_env/children/install_dep.sh
bash ${PWD}/scripts/install_dev_env/children/install_api_common_protos.sh
bash ${PWD}/scripts/install_dev_env/children/install_protoc.sh
bash ${PWD}/scripts/install_dev_env/children/install_proto_plugins.sh
bash ${PWD}/scripts/install_dev_env/children/install_prototool.sh
bash ${PWD}/scripts/install_dev_env/children/install_python_deps.sh
bash ${PWD}/scripts/install_dev_env/children/install_dart.sh