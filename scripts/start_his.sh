#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Start HIS docker container."
if [ $(docker images |grep his |wc -l) -ge '1' ];then
    docker run -d -p 8000:8000 --name hisv1 his > /dev/null
    echo "HIS docker container started!"
else
    echo "${RED}HIS docker container doesn't start, probably you have to generate image!${NC}"
fi
