#!/usr/bin/env bash
cd /opt/gopath/src/github.com/hyperledger/fabric-sdk-go/test/fixtures
docker-compose -f docker-compose.yaml stop
docker ps