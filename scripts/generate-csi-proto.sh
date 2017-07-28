#!/bin/bash
set -ex

pushd $(dirname $0)/../src/github.com/container-storage-interface/spec
  make csi.proto
  make csi.pb.go
popd

