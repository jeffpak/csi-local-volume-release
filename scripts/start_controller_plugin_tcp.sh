#!/bin/bash

set -x

cd `dirname $0`

pkill -f csi_local_controller

mkdir -p ~/csi_plugins
rm ~/csi_plugins/csi_local_controller.*

~/csi_local_controller &

