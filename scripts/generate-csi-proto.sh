#!/bin/bash
set -ex

pushd $(dirname $0)/../src/github.com/container-storage-interface/spec
  go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
  make csi.proto
  make csi.pb.go
popd

