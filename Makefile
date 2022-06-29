
run: generate
	go run cmd/foo/main.go

generate:
	go generate ./...
