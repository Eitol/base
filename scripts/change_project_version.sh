source ./utils.sh

OUT=run.bin
VERSION=`git describe --always --long --dirty`
BUILD_DATE_STR=`date +%Y-%m-%d\ %H:%M`
BUILD_DATE_TS=`date +%s`
VERSION_FILE=../config/version.go
MODULE_NAME=`get_current_module_name`


rm -f ${VERSION_FILE}
echo "package config" > ${VERSION_FILE}
echo "const (" >> ${VERSION_FILE}
echo "  Version = \"1.0\"" >> ${VERSION_FILE}
echo "  Compilation = \"${VERSION}\"" >> ${VERSION_FILE}
echo "  BuildDateStr = \"${BUILD_DATE_STR}\"" >> ${VERSION_FILE}
echo "  BuildDateTs = ${BUILD_DATE_TS}" >> ${VERSION_FILE}
echo "  ProjectName = \"${MODULE_NAME}\"" >> ${VERSION_FILE}
echo ")" >> ${VERSION_FILE}