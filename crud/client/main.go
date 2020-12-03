package main

import (
	"context"
	crudapi "crud/proto"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	client := crudapi.NewBlogServoceClient(conn)

	r := crudapi.CreateBlogRequest{
		Blog: &crudapi.Blog{
			AuthorId: "hitler",
			Title:    "meincampf",
			Content:  "minecraft",
		},
	}

	resp1, err := client.CreateBlog(context.Background(), &r)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("CreateBlog response", resp1.GetBlog())

	//read

	fmt.Println("reading the blog")

	resp2, err := client.ReadBlog(context.Background(),
		&crudapi.ReadBlogRequest{
			BlogId: "test1212",
		})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Println(respErr.Message())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("invalid argument")
			}
		}
	}

	blogid := resp1.GetBlog().GetId()

	resp2, err = client.ReadBlog(context.Background(),
		&crudapi.ReadBlogRequest{
			BlogId: blogid,
		})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Println(respErr.Message())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("invalid argument")
			}
		}
	}

	fmt.Println("ReadBlog response", resp2.GetBlog())

	// update

	fmt.Println("updating the blog")

	r3 := crudapi.UpdateBlogRequest{
		Blog: &crudapi.Blog{
			Id:       blogid,
			AuthorId: "changed",
			Title:    "another title",
			Content:  "different content",
		},
	}

	resp3, err := client.UpdateBlog(context.Background(), &r3)
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Println(respErr.Message())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("invalid argument")
			}
		}
	}
	fmt.Println("UpdateBlog response", resp3.GetBlog())

	// delete

	fmt.Println("DeleteBlog response")

	resp4, err := client.DeleteBlog(context.Background(), &crudapi.DeleteBlogRequest{
		BlogId: blogid,
	})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Println(respErr.Message())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("invalid argument")
			}
		}
	}
	fmt.Println("DeleteBlog response", resp4.GetBlogId())

	// list

	fmt.Println("lising all existing blogs")

	stream, err := client.ListBlog(context.Background(), &crudapi.ListBlogRequest{})
	if err != nil {
		log.Fatalf("error while calling ListBlog %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v", err)
		}
		fmt.Println(res.GetBlog())
	}
}
