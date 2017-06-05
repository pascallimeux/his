#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Generate keys for HIS ($PROJECTPATH/keys/server.key  and $PROJECTPATH/keys/server.pem)"
if [ -f $PROJECTPATH/keys/server.key ]; then
    echo "remove old server.key"
    rm $PROJECTPATH/keys/server.key
fi
if [ -f $PROJECTPATH/keys/server.crt ]; then
    echo "remove old server.crt"
    rm $PROJECTPATH/keys/server.crt
fi
openssl genrsa -out $PROJECTPATH/keys/server.key 2048 > /dev/null 2>&1
openssl req -new -x509 -sha256 -key $PROJECTPATH/keys/server.key -days 3650 -subj "/C=FR/ST=France/L=Grenoble/O=Orange/OU=OLS/CN=orange-labs.fr" -out $PROJECTPATH/keys/server.crt > /dev/null
if [ -f $PROJECTPATH/keys/server.crt  -o  -f $PROJECTPATH/keys/server.key ]; then
    echo "keys produced!"
fi