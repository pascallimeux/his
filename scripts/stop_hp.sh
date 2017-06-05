#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Stop Hyperledger test environment."
cd $GOPATH/src/github.com/hyperledger/fabric-sdk-go/test/fixtures && docker-compose -f docker-compose.yaml stop > /dev/null 2>&1
if [ $(docker ps |grep fabric |wc -l) -eq '0' ];then
    echo "  Environement stopped!"
else
    echo "${RED}Error environment not stopped!${NC}"
fi