#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Update configuration file ($PROJECTHISPATH/fixtures/config/config_prod.yaml)"
cp $PROJECTHISPATH/fixtures/config/config.yaml $PROJECTHISPATH/fixtures/config/config_prod.yaml
sudo sed -i "s/IPPEER0/$IPPEER0/g"  $PROJECTHISPATH/fixtures/config/config_prod.yaml
sudo sed -i "s/PORTPEER0/$PORTPEER0/g"  $PROJECTHISPATH/fixtures/config/config_prod.yaml
sudo sed -i "s/PORTEVT0/$PORTEVT0/g"  $PROJECTHISPATH/fixtures/config/config_prod.yaml
sudo sed -i "s/IPPEER1/$IPPEER1/g"  $PROJECTHISPATH/fixtures/config/config_prod.yaml
sudo sed -i "s/PORTPEER1/$PORTPEER1/g"  $PROJECTHISPATH/fixtures/config/config_prod.yaml
sudo sed -i "s/PORTEVT1/$PORTEVT1/g"  $PROJECTHISPATH/fixtures/config/config_prod.yaml
sudo sed -i "s/IPORDERER0/$IPORDERER0/g"  $PROJECTHISPATH/fixtures/config/config_prod.yaml
sudo sed -i "s/PORTORDERER0/$PORTORDERER0/g"  $PROJECTHISPATH/fixtures/config/config_prod.yaml
sudo sed -i "s/IPCA0/$IPCA0/g"  $PROJECTHISPATH/fixtures/config/config_prod.yaml
sudo sed -i "s/PORTCA0/$PORTCA0/g"  $PROJECTHISPATH/fixtures/config/config_prod.yaml
echo "peer0:  $IPPEER0:$PORTPEER0"
echo "peer1:  $IPPEER1:$PORTPEER1"
echo "oderer: $IPORDERER0:$PORTORDERER0"
echo "ca:     $IPCA0:$PORTCA0"