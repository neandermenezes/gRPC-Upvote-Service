package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	pb "github.com/neandermenezes/gRPC-Upvote-Service/proto/pb"
	"github.com/neandermenezes/gRPC-Upvote-Service/server/controller"
	"github.com/neandermenezes/gRPC-Upvote-Service/server/repository"
	"github.com/neandermenezes/gRPC-Upvote-Service/server/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var (
	postRepository repository.PostRepository = repository.NewPostRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostServer     = controller.NewPostController(postService)
)

// aaa

func main() {
	ctx := context.Background()

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("No mongo uri set in env variables")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	repository.Collection = client.Database("upvote-db").Collection("posts")
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal("Failed to listen on 50051")
	}

	log.Println("Listening on port 50051")

	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, &controller.PostServer{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
