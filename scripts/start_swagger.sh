#!/usr/bin/env bash
. ./scripts/config.sh
#swagger serve --port=3000 --host=127.0.0.1 swagger.json --base-path=/swagger-ui
#go run $PROJECTPATH/cmd/his-server/main.go --host=127.0.0.1 --port=3000 &
echo "> Start swagger."
CMD="$PROJECTHISPATH/cmd/his-server/swagger --host=127.0.0.1 --port=3000"
eval "$CMD > /dev/null 2>&1 &"
PID1=`pidof swagger`
if [ -n "$PID1" ];then
    echo "swagger started, use http://127.0.0.1:3000/docs"
else
    echo "${RED}swagger is not started, probably it doesn't build!${NC}"
fi