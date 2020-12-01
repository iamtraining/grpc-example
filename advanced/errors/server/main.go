package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"

	root "error/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct{}

func (s *Server) Sqrt(ctx context.Context, r *root.Request) (*root.Response, error) {
	num := r.GetNum()

	if num < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("cannot root out of a negative number %d", num))
	}

	resp := root.Response{
		Result: math.Sqrt(float64(num)),
	}

	return &resp, nil
}

func main() {
	fmt.Println("Server started")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer()
	root.RegisterRootSquareServer(srv, &Server{})

	srv.Serve(listener)
}
