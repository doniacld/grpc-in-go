# Mood tracker using Go and gRPC

This project aims to build a mood tracker using Go and gRPC.

## Prerequisites

Install the latest stable version of [Go](https://go.dev/doc/install)

Install all dependencies by running 

    $ make init

or install them manually like below.

Install protoc dependencies to generate the client and the server in Go from the proto file

    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

Install [grpcurl](https://github.com/fullstorydev/grpcurl) to test easily gRPC request

    $ go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

## Run

Run the project, it will run on port `9000` by default.

    $ make run

Hit `Ctrl-C` to kill it.

### Example of request

    > Getting the date in unix time : $ date +'%s'

    $ grpcurl --v -plaintext \
    -proto cmd/api/protos/tracker-api.proto \
    -d '{"Date": "1641758989", "Mood": {"State":"Happy"}}' \
    localhost:9000 tracker.TrackerAPI.CreateLog

## Resources

- [gRPC standard methods: Create](https://google.aip.dev/133)
