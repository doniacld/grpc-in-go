# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOINSTALL=$(GOCMD) install
GOMOD=$(GOCMD) mod
GOGENERATE=$(GOCMD) generate
GOLIST=$(GOCMD) list

# Sources parameters
SOURCE_ENTRYPOINT=./cmd/main.go

# Binary parameters
BINARY_NAME=moodtracker
BINARY_DESTINATION=./bin
BINARY_PATH=$(BINARY_DESTINATION)/$(BINARY_NAME)

# Targets
init:
		$(GOINSTALL) google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
        $(GOINSTALL) google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
        $(GOINSTALL) github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
build:
		$(GOBUILD) -o $(BINARY_PATH) -v $(SOURCE_ENTRYPOINT)
run:
		$(GORUN) $(SOURCE_ENTRYPOINT)
test: 
		$(GOTEST) -v ./...
clean: 
		$(GOCLEAN) $(SOURCE_ENTRYPOINT)
		rm -f $(BINARY_PATH)
proto:
		protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative cmd/api/protos/*.proto
