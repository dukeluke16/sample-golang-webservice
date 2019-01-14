#!/bin/sh
set -e

# Delete release Binary
rm -rf /go/bin/*

# Git Clean
git clean -xdf
