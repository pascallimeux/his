#!/usr/bin/env bash
. ./scripts/config.sh
echo "> Generate keys for HIS ($PROJECTHISPATH/keys/server.key  and $PROJECTHISPATH/keys/server.pem)"
if [ -f $PROJECTHISPATH/keys/server.key ]; then
    echo "remove old server.key"
    rm $PROJECTHISPATH/keys/server.key
fi
if [ -f $PROJECTHISPATH/keys/server.crt ]; then
    echo "remove old server.crt"
    rm $PROJECTHISPATH/keys/server.crt
fi
if [ ! -d $PROJECTHISPATH/keys ]; then
    mkdir $PROJECTHISPATH/keys
fi
openssl genrsa -out $PROJECTHISPATH/keys/server.key 2048 > /dev/null 2>&1
openssl req -new -x509 -sha256 -key $PROJECTHISPATH/keys/server.key -days 3650 -subj "/C=FR/ST=France/L=Grenoble/O=Orange/OU=OLS/CN=orange-labs.fr" -out $PROJECTHISPATH/keys/server.crt > /dev/null
if [ -f $PROJECTHISPATH/keys/server.crt  -o  -f $PROJECTHISPATH/keys/server.key ]; then
    echo "keys produced!"
fi
