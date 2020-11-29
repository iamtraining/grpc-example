package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	maxapi "bidi/proto"
)

type Server struct{}

func (s *Server) Max(stream maxapi.MaxService_MaxServer) error {
	var prev, max int32

	for {
		r, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalln(err)
			return err
		}

		prev = r.GetNum()
		if max == 0 {
			max = r.GetNum()
			if err := stream.Send(&maxapi.MaxResponse{Max: max}); err != nil {
				log.Fatalln(err)
				return err
			}
		}

		if prev > max {
			max = prev
			if err := stream.Send(&maxapi.MaxResponse{Max: max}); err != nil {
				log.Fatalln(err)
				return err
			}
		}
	}
}

func main() {
	fmt.Println("Server started")

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer()
	maxapi.RegisterMaxServiceServer(srv, &Server{})

	if err := srv.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
