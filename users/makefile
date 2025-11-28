.PHONY: proto clean deps run test

proto:
	@echo "Generating protobuf code..."
	protoc --go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		service/user.proto
	@echo "✅ Proto files generated successfully!"

clean:
	@echo "Cleaning generated files..."
	rm -rf pb/users
	@echo "✅ Cleaned!"

deps:
	@echo "Installing Go dependencies..."
	go get -u google.golang.org/grpc
	go get -u google.golang.org/protobuf
	@echo "✅ Dependencies installed!"

run:
	@echo "🚀 Starting Users gRPC Server..."
	go run cmd/server/main.go

test:
	@echo "Running tests..."
	go test -v ./...
	@echo "✅ Tests completed!"
