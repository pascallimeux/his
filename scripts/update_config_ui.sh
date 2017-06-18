#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Update configuration file ($PROJECTUIPATH/src/config.json)"
cp $PROJECTUIPATH/src/config.tmpl $PROJECTUIPATH/src/config.json
sudo sed -i "s/OCMSURL/$OCMSURL/g"  $PROJECTUIPATH/src/config.json
echo "ocms url:  $OCMSURL"
