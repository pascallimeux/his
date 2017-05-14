#!/usr/bin/env bash
cd $GOPATH/src/github.com/pascallimeux/his
echo "Build HIS..."
go build his.go
docker build -t his .
docker images
