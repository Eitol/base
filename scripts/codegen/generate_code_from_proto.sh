#!/usr/bin/env bash
CDIR=`dirname $0`
PROTO_GEN_OUT="api/gen"
PROTO_PKG="api/protos"

rm -rf ${PROTO_GEN_OUT}
prototool all --fix api/protos
python3 ${CDIR}/python_gen_fixer.py "${PROTO_GEN_OUT}/python" ${PROTO_PKG}