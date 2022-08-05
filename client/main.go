package main

import (
	pb "github.com/neandermenezes/gRPC-Upvote-Service/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewPostServiceClient(conn)

	id := createPost(c)
	readPost(c, id)
	updatePost(c, id)
	upvotePost(c, id)
	listPosts(c)
}
