#! /usr/bin/env bash

set -euo pipefail

mkdir -p "$PWD"/go/pkg/mod
export GOPATH=$PWD/go
export GOBIN=$GOPATH/bin
export PATH=$GOBIN:$PATH
