dc up:
	docker-compose up

loadenv:
	source ~/.bash_profile

protoc:
	protoc --proto_path=./proto/ --go_out=./ --go-grpc_out=. proto/token.proto		

build:
	CGO_ENABLED=0 GOOS=darwin  go build -o shrine ./cmd/shrine

run:
	go run ./cmd/shrine/*

test:
	go test ./...
