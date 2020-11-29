package main

import (
	bidirect "bi/proto"
	"context"
	"fmt"
	"io"
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

	client := bidirect.NewGreetEveryoneServiceClient(conn)

	wait := make(chan struct{})

	stream, err := client.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	names := []string{"DEN", "DENNI", "DENNY", "DENIS", "DENISQA"}

	requests := make([]*bidirect.GreetEveryoneRequest, len(names))

	for i, v := range names {
		requests[i] = &bidirect.GreetEveryoneRequest{
			Greet: &bidirect.Greeting{
				FirstName: v,
			},
		}
	}

	go func() {
		for _, r := range requests {
			if err := stream.Send(r); err != nil {
				log.Fatalln(err)
			}
			time.Sleep(1 * time.Second)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			resp, err := stream.Recv()

			if err == io.EOF {
				log.Println(err)
				break
			}

			if err != nil {
				log.Fatalln(err)
				break
			}

			fmt.Println(resp.GetResult())
		}

		close(wait)
	}()

	<-wait
}
