#!/bin/bash
# vim: set ft=sh

set -e

wget https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip
unzip protoc-3.3.0-linux-x86_64.zip
mv bin/protoc /usr/bin

csi-local-volume-release/scripts/generate-csi-proto.sh
cd csi-local-volume-release/

export GOROOT=/usr/local/go
export PATH=$GOROOT/bin:$PATH

export GOPATH=$PWD
export PATH=$PWD/bin:$PATH

go get github.com/onsi/ginkgo/ginkgo

./scripts/run_csi_cert.sh
