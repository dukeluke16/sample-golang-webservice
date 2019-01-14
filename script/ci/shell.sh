#!/bin/sh
set -e

if [ -z $WORKSPACE ]
then
  WORKSPACE=$(pwd)
fi

docker pull dukeluke16/golang-build:latest

docker run -i --rm \
  --entrypoint /bin/sh \
  -v $WORKSPACE:/go/src/github.com/dukeluke16/sample-golang-webservice \
  -v $WORKSPACE/release:/go/bin \
  -w /go/src/github.com/dukeluke16/sample-golang-webservice \
  -p 4001:4001 \
  dukeluke16/golang-build:latest $*
