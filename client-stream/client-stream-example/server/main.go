package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	greetpkg "client-stream-example/proto"
)

type Server struct{}

func (s *Server) LongGreet(stream greetpkg.LongGreetService_LongGreetServer) error {
	result := ""
	for {
		r, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&greetpkg.LongGreetResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalln(err)
		}

		result += "hello " + r.GetGreeting().GetFirstName() + " "
	}
}

func main() {
	fmt.Println("Server started")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer()
	greetpkg.RegisterLongGreetServiceServer(srv, &Server{})

	if err := srv.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}
