package main

import (
	"context"
	root "error/proto"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Client started")

	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	client := root.NewRootSquareClient(conn)

	num := -2

	r := root.Request{
		Num: int32(num),
	}

	resp, err := client.Sqrt(context.Background(), &r)

	if err != nil {
		respErr, ok := status.FromError(err)
		// если ок == тру то это ошибка grpc
		if ok {
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())

			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("probably negative argument")
			}
			return
		} else {
			log.Fatalln(err)
		}
	}

	fmt.Printf("root of %d is %f", num, resp.GetResult())
}
