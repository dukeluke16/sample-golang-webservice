#!/bin/sh
set -e

export BINARY_VERSION=$1
if [ -z $BINARY_VERSION ]
then
  BINARY_VERSION=0.0.1
fi

echo "BINARY_VERSION: $BINARY_VERSION"

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o release/service -ldflags "-X github.com/dukeluke16/sample-golang-webservice/config.BinaryVersion=${BINARY_VERSION} -X github.com/dukeluke16/sample-golang-webservice/web.EnableNewRelic=true" .
