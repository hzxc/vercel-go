package grpc

import (
	"net/http"
	"vercel-go/app/pingpong"
	pingpongpb "vercel-go/gen/go/pingpong/v1"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

type Handler http.Handler

func New() Handler {
	s := grpc.NewServer()
	pingpongpb.RegisterPingPongServiceServer(s, &pingpong.Service{})
	return grpcweb.WrapServer(s, grpcweb.WithOriginFunc(func(origin string) bool {
		// Allow all origins, DO NOT do this in production
		return true
	}))
}
