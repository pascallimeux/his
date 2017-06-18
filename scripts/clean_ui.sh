#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Clean UI environment."
if [ -d $PROJECTUIPATH/build ]; then
    rm -Rf $PROJECTUIPATH/build > /dev/null
    echo "UI package removed!"
fi
if [ $(docker images |grep his-ui |wc -l) -eq '1' ];then
    docker rmi his-ui > /dev/null
    echo "docker his-ui docker image removed!"
fi
if [ -f $PROJECTUIPATH/build/docker-image/his-ui.tar ]; then
    rm $PROJECTUIPATH/build/docker-image/his-ui.tar > /dev/null
    echo "docker his-ui image tar removed!"
fi
