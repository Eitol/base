#!/usr/bin/env bash

if [[ -d ${PWD}/vendor/github.com/izumin5210 ]]; then
    rm -rf ${PWD}/vendor/github.com/izumin5210
fi

if [[ ! -d ${PWD}/vendor/github.com/izumin5210 ]]; then
    mkdir -p ${PWD}/vendor/github.com/izumin5210
fi

cd ${PWD}/vendor/github.com/izumin5210
git clone https://github.com/izumin5210/grapi