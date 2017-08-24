#!/bin/bash

# run cmd and return cmd's exit code
# $1 string cmd
# $2 int doNotExitWhenFail 
report(){
    echo "$1 ..."
    eval $1
    if [ $? -eq 0 ];then
        echo $1 'ok!'
        return 0
    else
        echo -e '\033[0;31;1m' $1 'fail!!!\033[0m'
        if [ ! -z "$2" ] && [ "$2" -eq 1 ]; then
            return 1
        fi
        exit 1
    fi
}