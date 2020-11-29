package main

import (
	maxapi "bidi/proto"
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client started")

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	client := maxapi.NewMaxServiceClient(conn)

	nums := []int32{4, 7, 2, 19, 4, 6}

	requests := make([]*maxapi.MaxRequest, len(nums))

	for i, v := range nums {
		requests[i] = &maxapi.MaxRequest{Num: v}
	}

	stream, err := client.Max(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	await := make(chan struct{})

	go func() {
		for _, r := range requests {
			if err := stream.Send(r); err != nil {
				log.Fatalln(err)
			}
		}

		if err := stream.CloseSend(); err != nil {
			log.Fatalln(err)
		}
	}()

	go func() {
		for {
			r, err := stream.Recv()

			if err == io.EOF {
				log.Println(err)
				break
			}

			if err != nil {
				log.Fatalln(err)
				break
			}

			fmt.Println("max is", r.GetMax())
		}

		close(await)
	}()

	<-await
}
