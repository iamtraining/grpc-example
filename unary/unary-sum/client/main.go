package main

import (
	"context"
	"fmt"
	"log"
	sumpkg "sum/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	client := sumpkg.NewSumServiceClient(conn)

	req := sumpkg.SumRequest{
		Nums: &sumpkg.Form{
			FirstTerm:  3,
			SecondTerm: 10,
		},
	}

	resp, err := client.Add(context.Background(), &req)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp.GetResult())
}
