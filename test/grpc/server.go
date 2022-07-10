package main

import (
	"log"
	"net"
	"vercel-go/app/pingpong"

	pingpongpb "vercel-go/gen/go/pingpong/v1"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8087")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pingpongpb.RegisterPingPongServiceServer(s, &pingpong.Service{})
	s.Serve(lis)
}
