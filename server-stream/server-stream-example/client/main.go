package main

import (
	"context"
	"fmt"
	greetpb "greetmanytimes/proto"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	client := greetpb.NewGreetServiceClient(conn)

	r := greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName:     "DEN",
			LastName:      "NIS",
			NumberOfTimes: 20,
		},
	}

	stream, err := client.GreetManyTimes(context.Background(), &r)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		resp, err := stream.Recv()

		fmt.Println(resp.GetResult())

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln(err)
		}

	}
}
