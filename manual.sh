#!/usr/bin/env bash

source /etc/profile
PACKAGE="github.com/mustafaerbay/cer"
VERSION="$(git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//')"
COMMIT_HASH="$(git rev-parse --short HEAD)"
BUILD_TIMESTAMP=$(date '+%Y-%m-%dT%H:%M:%S')

#LDFLAGS=(
#  "-X 'cer/config.BuildVersion=${VERSION}'" 
#  "-X 'cer/config.CommitHash=${COMMIT_HASH}'" 
#  "-X 'cer/config.BuildTime=${BUILD_TIMESTAMP}'"
#)

#go build -ldflags="${LDFLAGS[*]}" -o dist/cer_linux_386/cer
CGO_ENABLED=0 GOOS=linux GARCH=amd64 go build -ldflags="-X 'cer/config.BuildVersion=${VERSION}' -X 'cer/config.CommitHash=${COMMIT_HASH}' -X 'cer/config.BuildTime=${BUILD_TIMESTAMP}'" -o dist/cer_linux_386/cer
docker build -t anatolman/cer:"${VERSION}" . 
