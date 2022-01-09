# grpc-in-go

Build a mood tracker using Go and gRPC.

## Prerequisites

Install [Go](https://go.dev/doc/install) 1.17

Install protoc dependencies

    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

Install [grpcurl](https://github.com/fullstorydev/grpcurl)

    $ go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

## Run

Run the project

    $ make run

### Example of request

    > Getting the date in unix time : $ date +'%s'

    $ grpcurl --v -plaintext -import-path . \
    --import-path proto/third_party -proto cmd/api/protos/mood-api.proto \
    -d '{"State":"Happy", "Date": "1641758989"}' localhost:9000 mood.MoodAPI.AddMood
