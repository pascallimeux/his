#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Clean HIS environment."
if [ -f $PROJECTPATH/build/bin/his ]; then
    rm $PROJECTPATH/build/bin/his > /dev/null
    echo "HIS binary removed!"
fi
if [ $(docker images |grep his |wc -l) -eq '1' ];then
    docker rmi his > /dev/null
    echo "docker his docker image removed!"
fi
if [ -f $PROJECTPATH/build/image/his.tar ]; then
    rm $PROJECTPATH/build/image/his.tar > /dev/null
    echo "docker his image.tar removed!"
fi
if [ -f $PROJECTPATH/keys/server.crt ]; then
    rm $PROJECTPATH/keys/server* > /dev/null
    echo "Server keys removed!"
fi

