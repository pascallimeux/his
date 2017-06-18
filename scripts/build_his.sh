#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Build HIS binary ($PROJECTHISPATH/build/bin/his)"
if [ -f $PROJECTHISPATH/build/bin/his ]; then
    echo "remove old his binary!"
    rm $PROJECTHISPATH/build/bin/his > /dev/null
fi
go build -o $PROJECTHISPATH/build/bin/his $PROJECTHISPATH/his.go > /dev/null
if [ ! -d $PROJECTHISPATH/build/bin ];then
    mkdir -p $PROJECTHISPATH/build/bin
fi
if [ -f $PROJECTHISPATH/build/bin/his ]; then
    echo "HIS binary created!"
else
    echo "${RED}HIS binary not create!${NC}"
fi
