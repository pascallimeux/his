#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Save UI docker image ($PROJECTUIPATH/build/image/his-ui.tar)"
if [ ! -d $PROJECTUIPATH/build/docker-image ]; then
    mkdir $PROJECTUIPATH/build/docker-image
fi
docker save his-ui -o $PROJECTUIPATH/build/docker-image/his-ui.tar > /dev/null
if [ -f $PROJECTUIPATH/build/docker-image/his-ui.tar ]; then
    echo "file created!"
else
    echo "${RED}file not created!${NC}"
fi
# instruction to load and stard docker image of HIS
# docker load -i his-ui.tar
# docker run -d -p 8080:8080 --name hisui hisui
# docker exec -ti hisui bash
