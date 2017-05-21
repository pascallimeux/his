**** Get repos from github ****
    git clone https://gerrit.hyperledger.org/r/fabric-sdk-go
    git clone https://github.com/pascallimeux/his.git

*** Install madatories libs ***
    go get -u github.com/kardianos/govendor
    go get github.com/gorilla/mux
    go get github.com/op/go-logging

*** Update libs with vendor ***
    rm -R vendor
    govendor init
    govendor add +external

*** Start hyperledger containers for tests ***
    cd /opt/gopath/src/github.com/hyperledger/fabric-sdk-go/test/fixtures/
    docker-compose -f docker-compose.yaml up --force-recreate -d
    docker ps

*** Build and start HIS project ***
    cd /opt/gopath/src/github.com/pascallimeux/his
    go build his.go
    ./his

*** Stop hyperledger containers ***
    cd /opt/gopath/src/github.com/hyperledger/fabric-sdk-go/test/fixtures/
    docker-compose -f docker-compose.yaml stop

*** Clean environment ***
    sudo rm -R /tmp/
    remove containers
    docker ps -a
    sudo docker  rm $( docker ps -aq)

    remove images
    docker images
    sudo docker  images | awk '/vp|none|dev-/ { print $3}' | xargs docker rmi -f

    rm -fr /var/hyperledger/production/*
    rm -fr /home/blockchain/.fabric-ca-client/msp/

    rm /opt/gopath/src/github.com/hyperledger/fabric-sdk-go/test/fixtures/fabricca/tlsOrg1/fabric-ca-server.db


*** Create HIS docker image ***g
    go build his.go
    docker build -t his .

*** Verify image ***
    docker images

*** Start HIS container ***
    docker run -d -p 8000:8000 --name hisv1 his

*** Start a terminal in the HIS container ***
    docker exec -it hisv1 bash

*** Stop the HIS container ***
    docker kill hisv1

*** Delete HIS container ***
    docker rm hisv1

*** Delete HIS image ***
    docker rmi his

Create server certificates
    openssl genrsa -out server.key 2048
    openssl ecparam -genkey -name secp384r1 -out server.key
    openssl req -new -x509 -sha256 -key server.key -days 3650 -subj "/C=FR/ST=France/L=Grenoble/O=Orange/OU=OLS/CN=orange-labs.fr" -out server.crt


---------------------------------------------------------------------------------------------
Use swagger
link:  https://github.com/go-swagger/go-swagger/tree/master/fixtures/goparsing/petstore

Installation:
    go get -u github.com/go-swagger/go-swagger/cmd/swagger
    cd ../cmg/swagger
    go install swagger.go
    go get github.com/go-openapi/runtime
    go get github.com/tylerb/graceful
    go get github.com/jessevdk/go-flags
    go get golang.org/x/net/context

Add annotation in API
    in main.go: add //go:generate swagger generate spec
    in handlers: add annotions

Generate swagger.yml and swagger.json
    in ../github.com/pascallimeux/his
    swagger init spec \
      --title "H.I.S application" \
      --description "Hyperledger Interface Server" \
      --version 1.0.0 \
      --scheme http
    swagger generate spec -o ./swagger.json -i ./swagger.yml
    swagger validate ./swagger.json

update golang libs
    cd ../his
    go get -u -f ./...

Generate server and test it
    swagger generate server -f ./swagger.json
    cd ./cmd/his-server && go run main.go

Remove swagger
    cd his
    rm swagger* && sudo rm -R cmd && sudo rm -R restapi
---------------------------------------------------------------------------------------------



