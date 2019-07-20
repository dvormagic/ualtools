#!/bin/bash

set -eu

. infra/functions/functions.sh

run "go get -v -t -d ./..."
run "go test -v ./..."
run "mkdir -p workspace"
run "sed -i 's/dev/$(build-tag)/g' pkg/config/version.go"
run "go build -o workspace/ualtools ./cmd/ualtools"
