#!/bin/bash

set -eu

. infra/functions/functions.sh

GOOGLE_PROJECT=ual-tools

configure-google-cloud

run "sed -i 's/dev/$(build-tag)/g' pkg/config/version.go"
run 'ualtools go build -o ualtools ./cmd/ualtools'
run "gsutil -h 'Cache-Control: no-cache' cp ualtools gs://ualtools/bin/ualtools"

run "echo $(build-tag) > version"
run "gsutil -h 'Cache-Control: no-cache' cp version gs://ualtools/version-manifest/ualtools"

run "gsutil -h 'Cache-Control: no-cache' cp install/install.sh gs://ualtools/install/ualtools"

for FILE in containers/*/Dockerfile; do
  APP=$(basename $(dirname $FILE))
  docker-build-autotag eu.gcr.io/$GOOGLE_PROJECT/$APP containers/$APP/Dockerfile containers/$APP
done

git-tag
