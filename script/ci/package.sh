#!/bin/sh
set -e

export BINARY_VERSION=$1
if [ -z $BINARY_VERSION ]
then
  BINARY_VERSION=0.0.1
fi

echo "BINARY_VERSION: $BINARY_VERSION"

sed -i "s/.*LABEL Version=.*/LABEL Version=${BINARY_VERSION}/" Dockerfile
docker build -t ${BINARY_VERSION} .
