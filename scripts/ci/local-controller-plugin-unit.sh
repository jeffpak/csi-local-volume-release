#!/bin/bash

set -e -x

csi-local-volume-release/scripts/generate-csi-proto.sh
cp csi-spec/csi* csi-localvolume-release/src/github.com/jeffpak/csi
cd csi-local-volume-release/

export GOROOT=/usr/local/go
export PATH=$GOROOT/bin:$PATH

export GOPATH=$PWD
export PATH=$PWD/bin:$PATH

go install github.com/onsi/ginkgo/ginkgo

pushd src/github.com/jeffpak/local-controller-plugin
  ginkgo -r -keepGoing -p -trace -randomizeAllSpecs -progress --race "$@"
popd
