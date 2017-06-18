#!/usr/bin/env bash
. ./scripts/config.sh
echo "> UI docker image is building."
if [ $(docker images |grep his |wc -l) -eq '1' ];then
    docker rmi his-ui -f > /dev/null
    echo "remove old his-ui docker image!"
fi
docker build -t his-ui $PROJECTUIPATH > /dev/null
if [ $(docker images |grep his-ui |wc -l) -eq '1' ];then
    echo "UI docker image created!"
else
    echo "${RED}UI docker image not create!${NC}"
fi
