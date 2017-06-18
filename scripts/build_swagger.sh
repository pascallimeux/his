#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Swagger is building."
cd $PROJECTHISPATH
swagger init spec \
    --title "his" \
    --description "Hyperledger Interface Server" \
    --version 1.0.0 \
    --scheme http  > /dev/null 2>&1 
swagger generate spec -o ./swagger.json -i ./swagger.yml  > /dev/null  2>&1
go get -u -f ./... > /dev/null  2>&1
swagger generate server -f ./swagger.json -A his > /dev/null 2>&1
if ls $PROJECTHISPATH/swagger.* 1> /dev/null 2>&1; then
    echo "swagger files created (swagger.yml swagger.json)"
else
    echo "${RED}swagger files not created!${NC}"
fi
if [ -d $PROJECTHISPATH/cmd ]; then
    echo "project cmd directory created ($PROJECTHISPATH/cmd)"
    go build -o $PROJECTHISPATH/cmd/his-server/swagger $PROJECTHISPATH/cmd/his-server/main.go > /dev/null
else
    echo "${RED}project cmd directory not created!${NC}"
fi
if [ -d $PROJECTHISPATH/restapi ]; then
    echo "project restapi directory created ($PROJECTHISPATH/restapi), to ask api use: http://ip:3000/docs"
else
    echo "${RED}project restapi directory not created!${NC}"
fi
