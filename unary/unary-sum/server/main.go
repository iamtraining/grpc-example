package main

import (
	"context"
	"errors"
	"log"
	"net"

	sumpkg "sum/proto"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Add(ctx context.Context, r *sumpkg.SumRequest) (*sumpkg.SumResponse, error) {
	first := r.GetNums().GetFirstTerm()
	second := r.GetNums().GetSecondTerm()
	var err error

	if first == 0 && second == 0 {
		err = errors.New("both values ​​are 0. they may not be specified")
		return nil, err
	}

	result := first + second

	resp := sumpkg.SumResponse{
		Result: result,
	}

	return &resp, err
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer()
	sumpkg.RegisterSumServiceServer(srv, &server{})

	if err := srv.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}
