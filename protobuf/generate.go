package protobuf

//go:generate protoc --go_out=. --go_opt=paths=source_relative ./user.proto ./log.proto

//go:generate go run ../cmd/gen/main.go
