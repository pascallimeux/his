#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Update configuration file ($PROJECTPATH/fixtures/config/config_prod.yaml)"
cp $PROJECTPATH/fixtures/config/config.yaml $PROJECTPATH/fixtures/config/config_prod.yaml
sudo sed -i "s/IPPEER0/$IPPEER0/g"  $PROJECTPATH/fixtures/config/config_prod.yaml
sudo sed -i "s/PORTPEER0/$PORTPEER0/g"  $PROJECTPATH/fixtures/config/config_prod.yaml
sudo sed -i "s/PORTEVT0/$PORTEVT0/g"  $PROJECTPATH/fixtures/config/config_prod.yaml
sudo sed -i "s/IPPEER1/$IPPEER1/g"  $PROJECTPATH/fixtures/config/config_prod.yaml
sudo sed -i "s/PORTPEER1/$PORTPEER1/g"  $PROJECTPATH/fixtures/config/config_prod.yaml
sudo sed -i "s/PORTEVT1/$PORTEVT1/g"  $PROJECTPATH/fixtures/config/config_prod.yaml
sudo sed -i "s/IPORDERER0/$IPORDERER0/g"  $PROJECTPATH/fixtures/config/config_prod.yaml
sudo sed -i "s/PORTORDERER0/$PORTORDERER0/g"  $PROJECTPATH/fixtures/config/config_prod.yaml
sudo sed -i "s/IPCA0/$IPCA0/g"  $PROJECTPATH/fixtures/config/config_prod.yaml
sudo sed -i "s/PORTCA0/$PORTCA0/g"  $PROJECTPATH/fixtures/config/config_prod.yaml
echo "peer0:  $IPPEER0:$PORTPEER0"
echo "peer1:  $IPPEER1:$PORTPEER1"
echo "oderer: $IPORDERER0:$PORTORDERER0"
echo "ca:     $IPCA0:$PORTCA0"