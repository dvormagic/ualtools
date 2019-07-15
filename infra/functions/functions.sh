#!/bin/bash

set -eu


function run {
  echo
  echo " ✗ $1"
  bash -c "$1"
}

function docker-build-autotag {
  run "docker build -t container -f $2 $3"
  
  HASH=$(docker image inspect container -f '{{.Id}}' | cut -d ':' -f 2)
  VERSION=${HASH:0:12}

  run "docker tag container $1:latest"
  run "docker tag container $1:$VERSION"

  run "docker push $1:latest"
  run "docker push $1:$VERSION"
}

function build-tag {
  echo $(date +%Y%m%d).$BUILD_NUMBER
}

function git-tag {
  PREFIX=${1:-}
  if [ ! -z $PREFIX ]; then
    PREFIX="${PREFIX}-"
  fi
  TAG=$PREFIX$(build-tag)

  run "git tag -f $TAG"
  run "git push --force origin refs/tags/$TAG:refs/tags/$TAG"
}
