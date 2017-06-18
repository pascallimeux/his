#!/usr/bin/env bash
. ./scripts/config.sh
echo "> HIS docker image is building."
if [ $(docker images |grep his |wc -l) -eq '1' ];then
    docker rmi his -f > /dev/null
    echo "remove old his docker image!"
fi
docker build -t his $PROJECTHISPATH > /dev/null
if [ $(docker images |grep his |wc -l) -eq '1' ];then
    echo "HIS docker image created!"
else
    echo "${RED}HIS docker image not create!${NC}"
fi
