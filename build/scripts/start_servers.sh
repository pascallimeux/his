#!/usr/bin/env bash
CMD="./swagger  --host=0.0.0.0 --port=3000"
eval "$CMD > /dev/null 2>&1 &"
./his