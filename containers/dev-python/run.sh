#!/bin/bash

set -eu

if [[ -z $WORKDIR ]]; then
  APP=$SERVICE
else
  APP=$(basename $WORKDIR)
fi

cd /workspace/$WORKDIR

if [ -f "requirements.txt" ]; then
  pip install --trusted-host pypi.python.org -r requirements.txt
fi

python main.py
