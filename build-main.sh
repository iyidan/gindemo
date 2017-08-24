#!/bin/bash
WORKDIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd )"
cd $WORKDIR
echo "cd $WORKDIR ..."

. ./build-functions.sh

PROJECTNAME=${WORKDIR##*/}
PLATFORM=$1
VERSION=$2

echo "project: ${PROJECTNAME}, platform: ${PLATFORM}, version: ${VERSION}, build..."

if [ -z $PLATFORM ] || [ -z $VERSION ]; then
    echo "PLATFORM(arg1) or VERSION(arg2) empty!"
    exit 1
fi

cat > ./version.go <<EOF
package main

// VERSION current version of the program
// this code will generate by build-main.sh
const VERSION = "${VERSION}"

EOF

GOOS=${PLATFORM} GOARCH=amd64 go build -o "${PROJECTNAME}.${PLATFORM}"
