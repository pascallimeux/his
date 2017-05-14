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


*** Create HIS docker image ***
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