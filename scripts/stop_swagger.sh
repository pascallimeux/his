#!/usr/bin/env bash
PID1=`pidof swagger`
if [ -n "$PID1" ]
then
    echo "stopping swagger pid=$PID1"
    kill -9 $PID1 2>/dev/null
else
    echo "Could not stop swagger, probably it doesn't run:" >&2
fi