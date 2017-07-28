#!/bin/bash
set -ex

pushd $(dirname $0)/../src/github.com/container-storage-interface/spec
  go get -u github.com/golang/protobuf/protoc-gen-go
  go install github.com/golang/protobuf/proto
  make csi.proto
  make csi.pb.go
popd

