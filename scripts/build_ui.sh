#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Build UI to deploy ($PROJECTUIPATH/build)"
if [ -d $PROJECTUIPATH/build ]; then
    echo "remove old build repo"
    rm -Rf $PROJECTUIPATH/build > /dev/null 2>&1
fi
cd $PROJECTUIPATH && npm run build  > /dev/null 2>&1
if [ $? -eq 0 ]; then
    echo "npm build OK"
else
    echo "${RED}npm build fail!${NC}"
fi
tar cvzf ui.tar.gz $PROJECTUIPATH/build  > /dev/null 2>&1
if [ ! -d $PROJECTUIPATH/build/tar-build ]; then
    mkdir $PROJECTUIPATH/build/tar-build
fi
mv ui.tar.gz $PROJECTUIPATH/build/tar-build/
if [ -f $PROJECTUIPATH/build/tar-build/ui.tar.gz ]; then
    echo "UI tar gz created!"
else
    echo "${RED}UI tar gz not create!${NC}"
fi
