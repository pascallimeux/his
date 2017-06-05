#!/usr/bin/env bash
echo "> Get fabric-sdk-go."
if [ ! -d $GOPATH/src/github.com/hyperledger/fabric-sdk-go ]; then
    echo "git clone fabric-sdk-go"
    mkdir -p $GOPATH/src/github.com/hyperledger/fabric-sdk-go
    cd $GOPATH/src/github.com/hyperledger
    git clone http://gerrit.hyperledger.org/r/fabric-sdk-go > /dev/null 2>&1
else
    echo "fabric-sdk-go already downloaded"
fi