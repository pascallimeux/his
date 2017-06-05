#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Build HIS binary ($PROJECTPATH/build/bin/his)"
if [ -f $PROJECTPATH/build/bin/his ]; then
    echo "old HIS binary removed!"
    rm $PROJECTPATH/build/bin/his > /dev/null
fi
go build -o $PROJECTPATH/build/bin/his $PROJECTPATH/his.go > /dev/null
if [ -f $PROJECTPATH/build/bin/his ]; then
    echo "HIS binary created!"
else
    echo "${RED}HIS binary not create!${NC}"
fi