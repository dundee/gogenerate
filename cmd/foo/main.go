package main

import (
	"github.com/dundee/gogenerate/pkg/log"
	pb "github.com/dundee/gogenerate/protobuf"
)

func main() {
	// reading protobuf data from somewhere
	user := &pb.User{
		Id:   1,
		Name: "John Doe",
	}

	log.Info("Received user", user)
}
