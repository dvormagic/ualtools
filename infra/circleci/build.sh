#!/bin/bash

set -eu

. infra/functions/functions.sh

run "go get -v -t -d ./..."
run "go test -v ./..."

run "mkdir -p workspace"
run "sed -i 's/dev/$(build-tag)/g' pkg/config/version.go"

run "go build -o workspace/linux/ualtools ./cmd/ualtools"
run "env GOOS=darwin GOARCH=amd64 go build -o workspace/mac/ualtools ./cmd/ualtools"
run "env GOOS=windows GOARCH=amd64 go build -o workspace/windows/ualtools.exe ./cmd/ualtools"
