#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Clean Swagger generated code."
if ls $PROJECTPATH/swagger.* 1> /dev/null 2>&1; then
    rm $PROJECTPATH/swagger.* > /dev/null
    echo "swagger files removed!"
fi
if [ -d $PROJECTPATH/cmd ]; then
    rm -R $PROJECTPATH/cmd > /dev/null
    echo "project cmd directory removed!"
fi
if [ -d $PROJECTPATH/restapi ]; then
    rm -R $PROJECTPATH/restapi > /dev/null
    echo "project restapi directory removed!"
fi