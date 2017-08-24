#!/bin/bash
WORKDIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd )"
cd $WORKDIR
echo "cd $WORKDIR ..."

. ./build-functions.sh

PROJECTNAME=${WORKDIR##*/}

report "go clean"
report 'find ./ -name "*.fasthttp.gz" | xargs rm'
report "rm -rf ./runtime/*.log"
report "rm -rf ./runtime/logs/*"
report "rm -rf ./${PROJECTNAME}.*"
