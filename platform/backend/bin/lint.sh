#!/usr/bin/env bash

docker run --rm \
  --volume "${PWD}":/go/src/platform/backend \
  --volume "${PWD}/../../api":/go/src/api \
  --volume "${PWD}/../../go.work":/go/src/go.work \
  --workdir /go/src/platform/backend \
  golangci/golangci-lint:v2.5.0 \
  golangci-lint run "$@"
