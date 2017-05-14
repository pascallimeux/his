#!/usr/bin/env bash
sudo docker  rm $( docker ps -aq)
docker ps -a
sudo docker  images | awk '/vp|none|dev-/ { print $3}' | xargs docker rmi -f
docker images
sudo rm /opt/gopath/src/github.com/hyperledger/fabric-sdk-go/test/fixtures/fabricca/tlsOrg1/fabric-ca-server.db
sudo rm /opt/gopath/src/github.com/pascallimeux/his/fixtures/enroll_user/*.pem
sudo rm /opt/gopath/src/github.com/pascallimeux/his/fixtures/keystore/*_sk