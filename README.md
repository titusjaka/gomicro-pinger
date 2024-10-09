# gomicro-pinger

It's a stub service to check how go-micro server/client works.

## How to run

```bash
# run go-micro server
go run main.go ponger micro --port 8080

# run go-micro client
go run main.go pinger micro --port 8080

# run gRPC server
go run main.go ponger grpc --port 8080

# run gRPC client
go run main.go pinger grpc --port 8080

# run go-micro server with gRPC transport
go run main.go ponger micro-grpc --port 8080

# run go-micro client with gRPC transport
go run main.go pinger micro-grpc --port 8080
```