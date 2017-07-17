#!/usr/bin/env sh

et -e

cd csi-local-volume-release/

export GOROOT=/usr/local/go
export PATH=$GOROOT/bin:$PATH

export GOPATH=$PWD
export PATH=$PWD/bin:$PATH

pushd src/github.com/jeffpak/local-controller-plugin
  ginkgo -r -keepGoing -p -trace -randomizeAllSpecs -progress --race "$@"
popd
