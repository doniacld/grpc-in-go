package protos

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

// Server is the gRPC server holding endpoints implementation
type Server struct{}

// CreateLog adds a new mood.
func (s *Server) CreateLog(ctx context.Context, dl *DailyLog) (*emptypb.Empty, error) {
	log.Println("create log request")

	log.Printf("request date %d, mood %s", dl.Date, dl.Mood.State)

	return &emptypb.Empty{}, nil
}

func (s *Server) mustEmbedUnimplementedTrackerAPIServer() {
	//TODO implement me
}
