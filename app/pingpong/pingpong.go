package pingpong

import (
	"context"
	pingpongpb "vercel-go/gen/go/pingpong/v1"
)

type Service struct {
	pingpongpb.UnimplementedPingPongServiceServer
}

func (s *Service) PingPong(ctx context.Context, request *pingpongpb.PingPongRequest) (*pingpongpb.PingPongResponse, error) {
	return &pingpongpb.PingPongResponse{
		Pong: "pong",
	}, nil
}
