#!/usr/bin/env bash
cd /opt/gopath/src/github.com/hyperledger/fabric-sdk-go/test/fixtures
docker-compose -f docker-compose.yaml up --force-recreate -d
docker ps
docker run -d -p 8000:8000 --name hisv1 his