#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Build HIS binary ($PROJECTPATH/build/bin/his)"
if [ -f $PROJECTPATH/build/bin/his ]; then
    echo "remove old his binary!"
    rm $PROJECTPATH/build/bin/his > /dev/null
fi
go build -o $PROJECTPATH/build/bin/his $PROJECTPATH/his.go > /dev/null
if [ ! -d $PROJECTPATH/build/bin ];then
    mkdir -p $PROJECTPATH/build/bin
fi
if [ -f $PROJECTPATH/build/bin/his ]; then
    echo "HIS binary created!"
else
    echo "${RED}HIS binary not create!${NC}"
fi
