#!/bin/bash

wget https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip
unzip protoc-3.3.0-linux-x86_64.zip
mv bin/protoc /usr/bin

pushd csi-spec
  go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
  make csi.proto
  make csi.pb.go
popd

