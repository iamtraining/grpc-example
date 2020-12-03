package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	crudapi "crud/proto"
)

var collection *mongo.Collection

type Blog struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

type Server struct{}

func (s *Server) CreateBlog(ctx context.Context, r *crudapi.CreateBlogRequest) (*crudapi.CreateBlogResponse, error) {
	fmt.Println("CreateBlog request")

	blog := r.GetBlog()

	data := Blog{
		AuthorID: blog.GetAuthorId(),
		Content:  blog.GetContent(),
		Title:    blog.GetTitle(),
	}

	result, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("internal error %v", err),
		)
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("coudlnt convert to primitive.ObjectID %v", err),
		)
	}

	resp := crudapi.CreateBlogResponse{
		Blog: &crudapi.Blog{
			Id:       id.Hex(),
			AuthorId: data.AuthorID,
			Title:    data.Title,
			Content:  data.Content,
		},
	}

	return &resp, nil
}

func (s *Server) ReadBlog(ctx context.Context, r *crudapi.ReadBlogRequest) (*crudapi.ReadBlogResponse, error) {
	fmt.Println("ReadBlog request")

	blogID := r.GetBlogId()

	id, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "couldnt parse id")
	}

	data := Blog{}

	filter := bson.M{"_id": id}

	result := collection.FindOne(context.Background(), filter)

	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, "couldnt find a blog with specified id")
	}

	return &crudapi.ReadBlogResponse{
		Blog: &crudapi.Blog{
			Id:       data.ID.Hex(),
			AuthorId: data.AuthorID,
			Title:    data.Title,
			Content:  data.Content,
		},
	}, nil
}

func (s *Server) UpdateBlog(ctx context.Context, r *crudapi.UpdateBlogRequest) (*crudapi.UpdateBlogResponse, error) {
	fmt.Println("UpdateBlog request")

	blog := r.GetBlog()

	id, err := primitive.ObjectIDFromHex(blog.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "couldnt parse id")
	}

	data := Blog{}

	filter := bson.M{"_id": id}

	result := collection.FindOne(context.Background(), filter)
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.Internal, "decoding error")
	}

	data.AuthorID = blog.GetAuthorId()
	data.Title = blog.GetTitle()
	data.Content = blog.GetContent()

	_, err = collection.ReplaceOne(context.Background(), filter, data)
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.Internal, "updating error")
	}

	return &crudapi.UpdateBlogResponse{
		Blog: &crudapi.Blog{
			Id:       blog.GetId(),
			AuthorId: blog.GetAuthorId(),
			Title:    blog.GetTitle(),
			Content:  blog.GetContent(),
		},
	}, nil
}

func (s *Server) DeleteBlog(ctx context.Context, r *crudapi.DeleteBlogRequest) (*crudapi.DeleteBlogResponse, error) {
	fmt.Println("DeleteBlog request")

	id, err := primitive.ObjectIDFromHex(r.GetBlogId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "couldnt parse id")
	}

	filter := bson.M{"_id": id}

	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "deleting error")
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "couldnt find document in mongodb")
	}

	return &crudapi.DeleteBlogResponse{BlogId: r.GetBlogId()}, nil
}

func (s *Server) ListBlog(r *crudapi.ListBlogRequest, stream crudapi.BlogServoce_ListBlogServer) error {
	fmt.Println("ListBlog request")

	cursor, err := collection.Find(context.Background(), nil)
	if err != nil {
		return status.Errorf(codes.Internal,
			fmt.Sprintf("unknown err %v", err))
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		data := Blog{}
		err := cursor.Decode(&data)
		if err != nil {
			return status.Errorf(codes.Internal,
				fmt.Sprintf("error while decoding %v", err))
		}

		stream.Send(&crudapi.ListBlogResponse{
			Blog: &crudapi.Blog{
				Id:       data.ID.Hex(),
				AuthorId: data.AuthorID,
				Title:    data.Title,
				Content:  data.Content,
			},
		})

		if err := cursor.Err(); err != nil {
			if err != nil {
				return status.Errorf(codes.Internal,
					fmt.Sprintf("unknown err %v", err))
			}
		}
	}

	return nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("connecting to mongodb")

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:2717"))
	if err != nil {
		log.Fatalln(err)
	}

	if err := client.Connect(context.TODO()); err != nil {
		log.Fatalln(err)
	}

	collection = client.Database("testing").Collection("blog")

	fmt.Println("blod crud service started")
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer()
	crudapi.RegisterBlogServoceServer(srv, &Server{})

	go func() {
		fmt.Println("Server started")

		if err := srv.Serve(listener); err != nil {
			log.Fatalln(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch

	fmt.Println("stopping the server")
	srv.Stop()

	fmt.Println("closing the listener")
	listener.Close()

	fmt.Println("closing mongodb connection")
	client.Disconnect(context.TODO())

	fmt.Println("goodbye")
}
