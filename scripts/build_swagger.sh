#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Swagger is building."
cd $PROJECTPATH
swagger init spec \
    --title "his" \
    --description "Hyperledger Interface Server" \
    --version 1.0.0 \
    --scheme http  > /dev/null 2>&1
swagger generate spec -o ./swagger.json -i ./swagger.yml  > /dev/null 2>&1
go get -u -f ./...
swagger generate server -f ./swagger.json -A his > /dev/null 2>&1
if ls $PROJECTPATH/swagger.* 1> /dev/null 2>&1; then
    echo "swagger files created (swagger.yml swagger.json)"
else
    echo "${RED}swagger files not created!${NC}"
fi
if [ -d $PROJECTPATH/cmd ]; then
    echo "project cmd directory created ($PROJECTPATH/cmd)"
    go build -o $PROJECTPATH/cmd/his-server/swagger $PROJECTPATH/cmd/his-server/main.go > /dev/null
else
    echo "${RED}project cmd directory not created!${NC}"
fi
if [ -d $PROJECTPATH/restapi ]; then
    echo "project restapi directory created ($PROJECTPATH/restapi)"
else
    echo "${RED}project restapi directory not created!${NC}"
fi
