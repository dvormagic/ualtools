#!/bin/bash

set -eu

mkdir -p ~/bin
source ~/.bashrc

curl https://ualtools.io/bin/ualtools > ~/bin/ualtools
chmod +x ~/bin/ualtools

ualtools pull
