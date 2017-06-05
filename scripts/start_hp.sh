#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Start docker containers for Hyperledger test environment."
cd $GOPATH/src/github.com/hyperledger/fabric-sdk-go/test/fixtures && docker-compose -f docker-compose.yaml up --force-recreate -d > /dev/null 2>&1
if [ $(docker ps |grep fabric |wc -l) -eq '7' ];then
    echo "environement started!"
else
    echo "${RED}environment not started!${NC}"
fi