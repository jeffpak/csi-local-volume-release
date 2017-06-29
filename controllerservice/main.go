package main

import (
  "log"
  "net"
  "./models"

  "google.golang.org/grpc"
  "google.golang.org/grpc/reflection"
  "code.cloudfoundry.org/goshims/filepathshim"
  "code.cloudfoundry.org/goshims/osshim"
)

const (
  port = ":8999"
)


func main() {
  lis, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := grpc.NewServer()
  models.NewController(&osshim.OsShim{}, &filepathshim.FilepathShim{}, "mountPathPlaceholder")
  // Register reflection service on gRPC server.
  reflection.Register(s)
  if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}