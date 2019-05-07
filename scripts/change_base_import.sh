#!/usr/bin/env bash
#!/bin/bash

NEW_MODULE_NAME=$1

source ./utils.sh

replace_imports(){
    LATEST_PWD=`pwd`
    cd ${BASE_PATH}
    grep -rl "$1" * | while read f; do
      echo ${f}
      echo "s|$1|$2|g"
      sed -i "s|$1|$2|g" "${f}"
    done
    cd ${LATEST_PWD}
}

check_new_module_name(){
    if [[ ${NEW_MODULE_NAME} == "" ]]; then
        echo "ERROR: empty param 'new module name'"
        echo "Usage:   bash change_base_import.sh NEW_MODULE_NAME"
        echo "Example: bash change_base_import.sh github.com/CrazyUser/super_project"
        exit 1
    fi
}

check_new_module_name
module_name=`get_current_module_name`
echo "The current module name is \"${module_name}\""
echo "New module name: \"${NEW_MODULE_NAME}\""
replace_imports "\"${module_name}/" "\"${NEW_MODULE_NAME}/"
sed -i "s|${module_name}|${NEW_MODULE_NAME}|g" "${MOD_FILE}"