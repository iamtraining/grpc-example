package main

import (
	"context"
	decomp "decomp/proto"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	client := decomp.NewDecoServiceClient(conn)

	r := decomp.DecoRequest{
		Number: 120,
	}

	stream, err := client.Decomposite(context.Background(), &r)

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
