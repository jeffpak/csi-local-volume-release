package main

import (
  . "./models"
  . ".."

  "code.cloudfoundry.org/goshims/filepathshim"
  "code.cloudfoundry.org/goshims/osshim"
  "log"

  "google.golang.org/grpc"
  "golang.org/x/net/context"
  "fmt"
)


func main() {
  address := "localhost:50051"
  defaultName := "world"

  // Set up a connection to the server.
  conn, err := grpc.Dial(address, grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %v", err)
  }
  defer conn.Close()
  c := NewController(&osshim.OsShim{}, &filepathshim.FilepathShim{}, "mountPathPlaceholder")

  // Contact the server and print out its response.
  volumeName := defaultName

  c.CreateVolume(context.Background(), &CreateVolumeRequest{
    Version: &Version{},
    Name: &volumeName,
    VolumeCapabilities: []*VolumeCapability{{Value: &VolumeCapability_Mount{}}},
  })
  fmt.Println("main running")
}
