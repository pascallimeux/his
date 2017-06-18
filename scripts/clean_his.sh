#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Clean HIS environment."
if [ -d $PROJECTHISPATH/build ]; then
    rm -Rf $PROJECTHISPATH/build > /dev/null
    echo "his building repo removed!"
fi
if [ $(docker images |grep his |wc -l) -eq '1' ];then
    docker rmi his > /dev/null
    echo "his docker image removed!"
fi

if [ -f $PROJECTHISPATH/keys/server.crt ]; then
    rm $PROJECTHISPATH/keys/server* > /dev/null
    echo "Server keys for HIS removed!"
fi

