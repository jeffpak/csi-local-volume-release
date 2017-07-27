#!/bin/bash
# vim: set ft=sh

set -e

csi-local-volume-release/scripts/generate-csi-proto.sh
cd csi-local-volume-release/

export GOROOT=/usr/local/go
export PATH=$GOROOT/bin:$PATH

export GOPATH=$PWD
export PATH=$PWD/bin:$PATH

go get github.com/onsi/ginkgo/ginkgo

./scripts/run_csi_cert.sh
