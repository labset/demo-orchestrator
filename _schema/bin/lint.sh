#!/usr/bin/env bash

docker run --rm \
    --volume "${PWD}":/workspace/schema \
    --volume "${PWD}/../api":/workspace/api \
    --workdir /workspace/schema \
    bufbuild/buf lint --verbose