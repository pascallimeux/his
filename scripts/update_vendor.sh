#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Update vendor Install dependencies."
cd $PROJECTPATH
sudo rm -R $PROJECTPATH/vendor > /dev/null 2>&1
govendor init
govendor add +external