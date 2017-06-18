#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Save HIS docker image ($PROJECTHISPATH/build/image/his.tar)"
if [ ! -d $PROJECTHISPATH/build/docker-image ]; then
    mkdir $PROJECTHISPATH/build/docker-image
fi
docker save his -o $PROJECTHISPATH/build/docker-image/his.tar > /dev/null
if [ -f $PROJECTHISPATH/build/docker-image/his.tar ]; then
    echo "file created!"
else
    echo "${RED}file not created!${NC}"
fi
# instruction to load and stard docker image of HIS
# docker load -i his.tar
# docker run -d -p 8000:8000 --name hisv1 his
# docker exec -ti hisv1 bash
