package main

import (
	log "github.com/sirupsen/logrus"

	pb "github.com/dundee/gogenerate/protobuf"
)

func main() {
	// reading protobuf data from somewhere
	user := &pb.User{
		Id:   1,
		Name: "John Doe",
	}

	log.WithFields(log.Fields{"user.id": user.Id, "user.name": user.Name}).Info("Received user")
}
