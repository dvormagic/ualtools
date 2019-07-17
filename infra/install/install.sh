#!/bin/bash

set -eu

mkdir -p ~/bin
source ~/.bashrc

curl https://storage.googleapis.com/ualtools/bin/ualtools > ~/bin/ualtools
chmod +x ~/bin/ualtools

ualtools pull
