package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	bidirect "bi/proto"
)

type Server struct{}

func (s *Server) GreetEveryone(stream bidirect.GreetEveryoneService_GreetEveryoneServer) error {
	for {
		r, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalln(err)
			return err
		}

		first := r.GetGreet().GetFirstName()
		//last := r.GetGreet().GetLastName()

		result := "hello " + first + "! "

		if err := stream.Send(&bidirect.GreetEveryoneResponse{
			Result: result,
		}); err != nil {
			log.Fatalln(err)
			return err
		}
	}
}

func main() {
	fmt.Println("Server started")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer()
	bidirect.RegisterGreetEveryoneServiceServer(srv, &Server{})

	if err := srv.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}
