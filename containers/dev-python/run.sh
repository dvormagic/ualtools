#!/bin/bash

set -eu

if [[ -z $WORKDIR ]]; then
  APP=$SERVICE
else
  APP=$(basename $WORKDIR)
fi

cd /workspace/$WORKDIR

python main.py
