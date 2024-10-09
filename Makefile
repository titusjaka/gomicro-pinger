.PHONY: proto
proto:
	@protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=require_unimplemented_servers=false \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		--micro_out=. \
		proto/service.proto

.PHONY: build
build:
	@go build -o build/gomicro-pinger *.go
