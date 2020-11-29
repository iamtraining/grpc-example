package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	avgpkg "computeavg/proto"
)

type Server struct{}

func (s *Server) Avg(stream avgpkg.ComputeServise_AvgServer) error {
	var avg float64
	var sum int32
	var count int32

	for {
		r, err := stream.Recv()

		if err == io.EOF {
			avg = float64(sum) / float64(count)

			return stream.SendAndClose(&avgpkg.Response{
				Result: avg,
			})
		}

		if err != nil {
			log.Fatalln(err)
		}

		sum += r.GetNum()
		fmt.Println(r.GetNum())
		count++
	}

}

func main() {
	fmt.Println("Server started")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer()
	avgpkg.RegisterComputeServiseServer(srv, &Server{})

	if err := srv.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}
