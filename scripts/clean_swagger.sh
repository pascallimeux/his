#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Clean Swagger generated code."
if ls $PROJECTHISPATH/swagger.* 1> /dev/null 2>&1; then
    rm $PROJECTHISPATH/swagger.* > /dev/null
    echo "swagger files removed!"
fi
if [ -d $PROJECTHISPATH/cmd ]; then
    rm -R $PROJECTHISPATH/cmd > /dev/null
    echo "project cmd directory removed!"
fi
if [ -d $PROJECTHISPATH/restapi ]; then
    rm -R $PROJECTHISPATH/restapi > /dev/null
    echo "project restapi directory removed!"
fi