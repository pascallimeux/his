#!/usr/bin/env bash
echo "> Installing dependencies..."
go get -u github.com/kardianos/govendor
go get github.com/gorilla/mux
go get github.com/op/go-logging