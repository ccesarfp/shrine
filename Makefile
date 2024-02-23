compose up:
	docker-compose up

protoc:
	source ~/.bash_profile
	protoc --proto_path=./proto/ --go_out=./ --go-grpc_out=. proto/token.proto		

run:
	go run ./cmd/shrine/main.go

test:
	go test ./...
