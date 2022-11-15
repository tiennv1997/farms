#!/usr/bin/env bash

SCRIPT_PATH=${0%/*}
if [ "$0" != "$SCRIPT_PATH" ] && [ "$SCRIPT_PATH" != "" ]; then
    cd $SCRIPT_PATH
fi

abigen -abi ./abi/masterchefv2.abi -pkg contract -type MasterChefV2 -out ./generated/masterchefv2_abi.go