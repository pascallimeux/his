#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Clean Hyperledger test environment."
if [ $(docker ps -aq |wc -l) -ge '1' ];then
    sudo docker  rm $( docker ps -aq) > /dev/null
    if [ $? -eq 0 ];then
        echo "   All containers removed!"
    fi
fi
if [ $(sudo docker  images | awk '/vp|none|dev-/ { print $3}'|wc -l) -ge '1' ];then
    sudo docker  images | awk '/vp|none|dev-/ { print $3}' | xargs docker rmi -f > /dev/null
    if [ $? -eq 0 ];then
        echo "all smart contract containers removed!"
    fi
fi
if [ -f $GOPATH/src/github.com/hyperledger/fabric-sdk-go/test/fixtures/fabricca/tlsOrg1/fabric-ca-server.db ];then
    sudo rm $GOPATH/src/github.com/hyperledger/fabric-sdk-go/test/fixtures/fabricca/tlsOrg1/fabric-ca-server.db > /dev/null
    if [ $? -eq 0 ];then
        echo "fabric-sdk DB removed!"
    fi
fi
if [ -f $PROJECTPATH/fixtures/enroll_user/*.pem ];then
    sudo rm $PROJECTPATH/fixtures/enroll_user/*.pem > /dev/null
    if [ $? -eq 0 ];then
        echo "enroll pem files removed!"
    fi
fi
if [ -f $PROJECTPATH/fixtures/enroll_user/*.json ];then
    sudo rm $PROJECTPATH/fixtures/enroll_user/*.json > /dev/null
    if [ $? -eq 0 ];then
        echo "enroll json files removed!"
    fi
fi
if [ -f $PROJECTPATH/fixtures/keystore/*_sk ];then
    sudo rm $PROJECTPATH/fixtures/keystore/*_sk > /dev/null
    if [ $? -eq 0 ];then
    echo "stored keys removed!"
    fi
fi
return 0