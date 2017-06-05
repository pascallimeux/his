#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Update fabric-sdk-go."
if [ -d $GOPATH/src/github.com/hyperledger/fabric-sdk-go ]; then
    echo "git pull fabric-sdk-go"
    cd $GOPATH/src/github.com/hyperledger/fabric-sdk-go
    git pull
else
    echo "git clone fabric-sdk-go"
    mkdir -p $GOPATH/src/github.com/hyperledger/fabric-sdk-go
    cd $GOPATH/src/github.com/hyperledger
    git clone http://gerrit.hyperledger.org/r/fabric-sdk-go > /dev/null 2>&1
fi