#!/bin/bash
# vim: set ft=sh

set -e

cp csi-spec/created-csi-proto* csi-localvolume-release/src/github.com/jeffpak/csi
cd csi-local-volume-release/

export GOROOT=/usr/local/go
export PATH=$GOROOT/bin:$PATH

export GOPATH=$PWD
export PATH=$PWD/bin:$PATH

go get github.com/onsi/ginkgo/ginkgo

./scripts/run_csi_cert.sh
