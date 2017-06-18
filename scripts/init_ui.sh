#!/usr/bin/env bash
. ./scripts/config.sh
echo "> npm install for his-ui"
cd $PROJECTUIPATH && sudo npm install
