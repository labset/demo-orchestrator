#!/usr/bin/env bash

docker run --rm -v $(pwd):/src -w /src sqlc/sqlc:1.29.0 generate
