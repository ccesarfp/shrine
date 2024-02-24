compose up:
	docker-compose up

loadenv:
	source ~/.bash_profile

protoc:
	protoc --proto_path=./proto/ --go_out=./ --go-grpc_out=. proto/token.proto		

run:
	go run ./cmd/shrine/main.go

test:
	go test ./...

cli:
	go run ./cmd/shrine-cli/*
