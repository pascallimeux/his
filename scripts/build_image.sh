#!/usr/bin/env bash
. ./scripts/config.sh
echo "> HIS docker image is building."
if [ $(docker images |grep his |wc -l) -eq '1' ];then
    docker rmi his > /dev/null
    echo "old HIS docker image removed!"
fi
docker build -t his $PROJECTPATH > /dev/null
if [ $(docker images |grep his |wc -l) -eq '1' ];then
    echo "HID docker image created!"
else
    echo "${RED}HIS docker image not create!${NC}"
fi
