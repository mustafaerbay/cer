#!/usr/bin/env bash

PACKAGE="github.com/mustafaerbay/cer"
VERSION="$(git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//')"
COMMIT_HASH="$(git rev-parse --short HEAD)"
BUILD_TIMESTAMP=$(date '+%Y-%m-%dT%H:%M:%S')

LDFLAGS=(
  "-X 'cer/config.BuildVersion=${VERSION}'"
  "-X 'cer/config.CommitHash=${COMMIT_HASH}'"
  "-X 'cer/config.BuildTime=${BUILD_TIMESTAMP}'"
)

go build -ldflags="${LDFLAGS[*]}" -o dist/cer_linux_386/cer

docker build .