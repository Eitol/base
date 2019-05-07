#!/usr/bin/env bash


BASE_PATH=../
MOD_FILE=${BASE_PATH}/go.mod

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