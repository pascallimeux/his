#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Start docker containers for his and ui test environment."
cd $PROJECTPATH && docker-compose -f docker-compose.yaml up --force-recreate -d
if [ $(docker ps |grep his |wc -l) -eq '2' ];then
    echo "environement started!"
else
    echo "${RED}environment not started!${NC}"
fi
