package main

import (
	avgpkg "computeavg/proto"
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
		log.Fatalln("1", err)
	}

	client := avgpkg.NewComputeServiseClient(conn)

	nums := []int32{3, 5, 9, 54, 23}

	requests := make([]*avgpkg.Request, len(nums))

	for i, num := range nums {
		requests[i] = &avgpkg.Request{
			Num: num,
		}
	}

	stream, err := client.Avg(context.Background())
	if err != nil {
		log.Fatalln("2", err)
	}

	for _, r := range requests {
		if err := stream.Send(r); err != nil {
			log.Fatalln("3", err)
		}
		time.Sleep(100 * time.Millisecond)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("4", err)
	}

	fmt.Println("the avg is", resp.GetResult())
}
