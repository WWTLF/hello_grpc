make:
	go build ./cmd/hello_server
	go build ./cmd/hello_client


proto:
	protoc --go_out=./pkg --go_opt=paths=source_relative \
    --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative \
    api/hello/*.proto

