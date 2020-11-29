package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"

	greetpb "greetmanytimes/proto"
)

type Server struct{}

func (s *Server) GreetManyTimes(r *greetpb.GreetRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	first := r.GetGreeting().GetFirstName()
	last := r.GetGreeting().GetLastName()
	n := r.GetGreeting().GetNumberOfTimes()

	for i := int32(1); i <= n; i++ {
		result := "hello " + first + " " + last + " " + strconv.Itoa(int(i))
		resp := greetpb.GreetResponse{
			Result: result,
		}
		if err := stream.Send(&resp); err != nil {
			log.Fatalln(err)
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

	grpcSrv := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(grpcSrv, &Server{})

	if err := grpcSrv.Serve(listener); err != nil {
		fmt.Println(err)
	}

}
