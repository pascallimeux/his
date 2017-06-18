#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Stop HIS and HIS-UI test environment."
cd $PROJECTPATH && docker-compose -f docker-compose.yaml stop > /dev/null 2>&1
if [ $(docker ps |grep his |wc -l) -eq '0' ];then
    echo "  Environement stopped!"
else
    echo "${RED}Error environment not stopped!${NC}"
fi
