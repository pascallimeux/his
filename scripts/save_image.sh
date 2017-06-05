#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Save HIS docker image ($PROJECTPATH/build/image/his.tar)"
docker save his -o $PROJECTPATH/build/image/his.tar > /dev/null
if [ -f $PROJECTPATH/build/image/his.tar ]; then
    echo "file created!"
else
    echo "${RED}file not created!${NC}"
fi
# instruction to load and stard docker image of HIS
# docker load -i his.tar
# docker run -d -p 8000:8000 --name hisv1 his
# docker exec -ti hisv1 bash