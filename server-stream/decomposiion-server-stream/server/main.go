package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	decomp "decomp/proto"
)

type Server struct{}

func (s *Server) Decomposite(r *decomp.DecoRequest, stream decomp.DecoService_DecompositeServer) error {
	num, k := r.GetNumber(), int32(2)

	for num > 1 {
		if num%k == 0 {
			resp := decomp.DecoResponse{
				Result: k,
			}

			if err := stream.Send(&resp); err != nil {
				log.Fatalln(err)
			}

			num /= k
		} else {
			k++
		}

		time.Sleep(1 * time.Second)
	}

	return nil
}

func main() {
	fmt.Println("Server started")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	decomp.RegisterDecoServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}
