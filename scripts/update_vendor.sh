#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Update vendor Install dependencies."
cd $PROJECTHISPATH
sudo rm -R $PROJECTHISPATH/vendor > /dev/null 2>&1
govendor init
govendor add +external