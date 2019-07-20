#!/bin/bash

set -eu

. infra/functions/functions.sh

run "gcloud auth activate-service-account --key-file=config/service-account.json"
run "gcloud --quiet config set project ${GOOGLE_PROJECT_ID}"
run "gcloud --quiet config set compute/zone ${GOOGLE_COMPUTE_ZONE}"
run "gcloud auth configure-docker"

run "gsutil -h 'Cache-Control: no-cache' cp workspace/ualtools gs://ualtools/bin/ualtools"

run "echo $(build-tag) > version"
run "gsutil -h 'Cache-Control: no-cache' cp version gs://ualtools/version-manifest/ualtools"

run "gsutil -h 'Cache-Control: no-cache' cp infra/install/install.sh gs://ualtools/install/ualtools"

for FILE in containers/*/Dockerfile; do
  APP=$(basename $(dirname $FILE))
  docker-build-autotag eu.gcr.io/$GOOGLE_PROJECT_ID/$APP containers/$APP/Dockerfile containers/$APP
done

git-tag
