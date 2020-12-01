package main

import (
	"context"
	greet "deadline/proto"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("client started")

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	client := greet.NewGreetServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()

	r := greet.GreetRequest{
		Greeting: &greet.Greeting{
			FirstName: "DEN",
			LastName:  "NIS",
		},
	}

	resp, err := client.Greet(ctx, &r)

	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			if respErr.Code() == codes.DeadlineExceeded {
				fmt.Println("deadline was exceeded", respErr.Message())
				return
			} else {
				log.Fatalln(err)
				return
			}
		}
	}

	fmt.Println(resp.GetResult())
}
