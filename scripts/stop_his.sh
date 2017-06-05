#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Stop HIS docker container."
if [ $(docker ps |grep his |wc -l) -ge '1' ];then
    docker kill hisv1 > /dev/null
    echo "his docker container stopped!"
else
    echo "${RED}HIS docker container doesn't stop, probably it doesn't run!${NC}"
fi
return 0