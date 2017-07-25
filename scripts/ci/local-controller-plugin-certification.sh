#!/bin/bash
# vim: set ft=sh

set -e

./scripts/generate-csi-proto.sh
cp csi-spec/csi* csi-local-volume-release/src/github.com/jeffpak/csi
cd csi-local-volume-release/

export GOROOT=/usr/local/go
export PATH=$GOROOT/bin:$PATH

export GOPATH=$PWD
export PATH=$PWD/bin:$PATH

go get github.com/onsi/ginkgo/ginkgo

./scripts/run_csi_cert.sh
