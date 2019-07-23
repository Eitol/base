#!/usr/bin/env bash
#!/bin/bash

NEW_MODULE_NAME=$1
BASE_PATH=./
MOD_FILE=${BASE_PATH}/go.mod

# CONST
INVALID_PARAM_VAL="invalid"
CALLER="make init new_name="

check_if_fail(){
   if [[ $? -eq 0 ]]; then
      echo "$1... OK"
   else
      echo "Fail in $1"
      exit 1
   fi
}

get_current_module_name(){
    if [[ ! -e ${MOD_FILE} ]]; then
        echo "ERROR: missing go.mod file"
        exit 1
    fi
    # Get the first line
    IN=$(awk 'FNR <= 1' ${MOD_FILE})
    # split the line in two. i.e: "module base" => ["module", "base"]
    arrIN=(${IN//;/ })
    len=${#arrIN[@]}
    if [[ ${arrIN[1]} == "" || ${len} -eq 0 ]]; then
        echo "ERROR: empty module name in go.mod file"
        exit 1
    fi
    echo ${arrIN[1]}
}

replace_imports(){
    LATEST_PWD=`pwd`
    CURRENT_MODULE_NAME=$1
    NEW_MODULE_NAME=$2
    cd ${BASE_PATH}
    grep -rl "${CURRENT_MODULE_NAME}" * | while read f; do
      echo "Edited file: ${f}"
      sed -i "s|$1|$2|g" "${f}"
    done
    cd ${LATEST_PWD}
}

check_new_module_name(){
    if [[ ${NEW_MODULE_NAME} == "" || ${NEW_MODULE_NAME} == ${INVALID_PARAM_VAL} ]]; then
        echo "ERROR: empty param 'new_module name'"
        echo "Usage:   ${CALLER}NEW_MODULE_NAME"
        echo "Example: ${CALLER}github.com/CrazyUser/super_project"
        exit 1
    fi
    if [[ "$(${NEW_MODULE_NAME} | grep /)" != "" ]]; then

    fi
}

check_new_module_name
check_if_fail "check_new_module_name"

module_name=`get_current_module_name`
echo "The current module name is \"${module_name}\""
echo "New module name: \"${NEW_MODULE_NAME}\""
replace_imports "\"${module_name}/" "\"${NEW_MODULE_NAME}/"
check_if_fail "replace_imports ${module_name} ${NEW_MODULE_NAME}"

sed -i "s|${module_name}|${NEW_MODULE_NAME}|g" "${MOD_FILE}"
check_if_fail "sed"