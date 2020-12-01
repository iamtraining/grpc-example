package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	greet "deadline/proto"
)

type Server struct{}

func (s *Server) Greet(ctx context.Context, r *greet.GreetRequest) (*greet.GreetResponse, error) {
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			return nil, status.Errorf(codes.Canceled, "the client canceled the request")
		}

		time.Sleep(1 * time.Second)
	}

	first := r.GetGreeting().GetFirstName()
	last := r.GetGreeting().GetLastName()

	resp := greet.GreetResponse{
		Result: fmt.Sprintf("Hello %s %s", first, last),
	}

	return &resp, nil
}

func main() {
	fmt.Println("Server started")

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer()

	greet.RegisterGreetServiceServer(srv, &Server{})

	srv.Serve(listen)
}
