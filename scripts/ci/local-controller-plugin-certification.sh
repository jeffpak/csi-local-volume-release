#!/bin/bash
# vim: set ft=sh

set -e

wget https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip
unzip protoc-3.3.0-linux-x86_64.zip
mv bin/protoc /usr/bin

work_dir=$(pwd)
script_dir=$(pwd)/csi-local-volume-release/scripts

cd csi-local-volume-release/

export GOROOT=/usr/local/go
export PATH=$GOROOT/bin:$PATH

export GOPATH=$PWD
export PATH=$PWD/bin:$PATH

go get github.com/onsi/ginkgo/ginkgo

pushd scripts
 ./generate-csi-proto.sh
popd

go build -o "$HOME/csi_local_controller" "src/github.com/jeffpak/local-controller-plugin/cmd/localcontrollerplugin/main.go"
go build -o "$HOME/csi_local_node" "src/github.com/jeffpak/local-node-plugin/cmd/localnodeplugin/main.go"

#=======================================================================================================================
# local-controller-plugin local-node-plugin
#=======================================================================================================================

function cleanup {
  cd $script_dir
  /bin/bash ./stop_controller_plugin_tcp.sh
  /bin/bash ./stop_node_plugin_tcp.sh
  rm -rf $HOME/csi_plugins
}
# TCP TESTS
export FIXTURE_FILENAME=${script_dir}/fixtures/local_plugin_cert.json
trap cleanup EXIT
/bin/bash scripts/start_controller_plugin_tcp.sh
/bin/bash scripts/start_node_plugin_tcp.sh

mkdir -p ${work_dir}/go/src/github.com/paulcwarren/

pushd ${work_dir}/go
  export GOROOT=/usr/local/go
  export PATH=$GOROOT/bin:$PATH
  export GOPATH=$PWD
  export PATH=$PWD/bin:$PATH
  ln -s ${work_dir}/csi-cert  src/github.com/paulcwarren/csi-cert
  cd src/github.com/paulcwarren/csi-cert
  ./scripts/go_get_all_dep.sh
  ginkgo -p
popd
