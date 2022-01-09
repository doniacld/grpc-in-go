package protos

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct{}

// AddMood adds a new mood.
func (s *Server) AddMood(ctx context.Context, mood *Mood) (*emptypb.Empty, error) {
	log.Println("add mood request")

	return &emptypb.Empty{}, nil
}

func (s *Server) mustEmbedUnimplementedMoodAPIServer() {
	//TODO implement me
}
