#!/bin/bash

set -eu

. infra/functions/functions.sh

run "gcloud auth configure-docker"

run "gsutil -h 'Cache-Control: no-cache' cp workspace/linux/ualtools gs://ualtools/linux/ualtools"
run "gsutil -h 'Cache-Control: no-cache' cp workspace/mac/ualtools gs://ualtools/mac/ualtools"
run "gsutil -h 'Cache-Control: no-cache' cp workspace/windows/ualtools.exe gs://ualtools/windows/ualtools.exe"

run "echo $(build-tag) > version"
run "gsutil -h 'Cache-Control: no-cache' cp version gs://ualtools/version-manifest/ualtools"

run "gsutil -h 'Cache-Control: no-cache' cp infra/install/linux-install.sh gs://ualtools/install/linux-ualtools"
run "gsutil -h 'Cache-Control: no-cache' cp infra/install/mac-install.sh gs://ualtools/install/mac-ualtools"

for FILE in containers/*/Dockerfile; do
  APP=$(basename $(dirname $FILE))
  docker-build-autotag eu.gcr.io/$GOOGLE_PROJECT_ID/$APP containers/$APP/Dockerfile containers/$APP
done

git-tag
