package main

import (
	greetpkg "client-stream-example/proto"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client started")

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	client := greetpkg.NewLongGreetServiceClient(conn)

	stream, err := client.LongGreet(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	requests := []*greetpkg.LongGreetRequest{}

	for _, s := range []string{"DENI", "DENNY", "DENIS", "DENNYS"} {
		requests = append(requests, &greetpkg.LongGreetRequest{
			Greeting: &greetpkg.Greeting{
				FirstName: s,
			},
		})
	}

	for _, r := range requests {
		if err := stream.Send(r); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(1 * time.Second)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp)
}
