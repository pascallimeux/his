#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Stop HIS docker container."
if [ $(docker ps |grep his |wc -l) -ge '1' ];then
    docker kill hisv1 > /dev/null
    echo "his docker container stopped!"
else
    echo "Could not stop HIS docker container, probably it doesn't run:!"
fi
return 0