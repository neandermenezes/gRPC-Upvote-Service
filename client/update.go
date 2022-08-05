package main

import (
	"context"
	pb "github.com/neandermenezes/gRPC-Upvote-Service/proto/pb"
	"log"
)

func updatePost(c pb.PostServiceClient, id string) {
	log.Println("---updatePost was invoked---")

	newBlog := &pb.Post{
		Id:         id,
		AuthorName: "Not Neander",
		Title:      "A new title",
		Content:    "New content",
	}

	_, err := c.UpdatePost(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Post was updated!")
}
