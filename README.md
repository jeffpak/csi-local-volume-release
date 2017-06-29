# csi-local-volume-release
Local Volume Release for Cloud Foundry that follows protocol specified by Container Storage Interface

Useful Commands:

To generate protobuf:
protoc -I . -I=/Users/pivotal/go/src/github.com/gogo/protobuf/protobuf -I=/Users/pivotal/go/src csi-local-volume-release.proto --go_out=plugins=grpc:.
