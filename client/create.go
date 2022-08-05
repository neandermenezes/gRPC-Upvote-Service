package main

import (
	"context"
	pb "github.com/neandermenezes/gRPC-Upvote-Service/proto/pb"
	"log"
)

func createPost(c pb.PostServiceClient) string {
	log.Println("---createPost was invoked---")

	blog := &pb.Post{
		AuthorName: "Neander",
		Title:      "New Post!",
		Content:    "GoLang gRPC is Amazing!",
	}

	res, err := c.CreatePost(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Post has been created: %s\n", res.String())
	return res.Id
}
