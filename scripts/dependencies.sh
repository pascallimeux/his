#!/usr/bin/env bash
echo "> Installing dependencies..."
go get -u github.com/kardianos/govendor
go get github.com/gorilla/mux
go get github.com/op/go-logging
go get  -u github.com/go-swagger/go-swagger/cmd/swagger
cd $GOPATH/src/github.com/go-swagger/go-swagger/cmd/swagger
go build swagger.go
sudo cp swagger /usr/local/bin
