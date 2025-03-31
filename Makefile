BINARY_NAME = ezex-notification
BUILD_DIR = build
CMD_DIR = internal/cmd/server/main.go


all: build test

########################################
### Tools needed for development
devtools:
	@echo "Installing devtools"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install mvdan.cc/gofumpt@latest

proto:
	protoc --go_out ./api/grpc/proto --go_opt paths=source_relative \
           --go-grpc_out ./api/grpc/proto --go-grpc_opt paths=source_relative \
	       --proto_path=./api/grpc/proto api/grpc/proto/*.proto

docker:
	docker build --tag ezex-notification .

########################################
### Building
build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)

release:
	@mkdir -p $(BUILD_DIR)
	go build -ldflags "-s -w" -trimpath -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)

clean:
	@echo "Cleaning up build artifacts..."
	rm -rf $(BUILD_DIR)

########################################
### Testing
test:
	go test ./...

########################################
### Formatting the code
fmt:
	gofumpt -l -w .

check:
	golangci-lint run --timeout=20m0s


.PHONY: devtools proto docker
.PHONY: build release
.PHONY: test
.PHONY: fmt check
