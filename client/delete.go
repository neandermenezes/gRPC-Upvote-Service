package main

import (
	"context"
	pb "github.com/neandermenezes/gRPC-Upvote-Service/proto/pb"
	"log"
)

func deletePost(c pb.PostServiceClient, id string) {
	log.Println("---deletePost was invoked---")
	_, err := c.DeletePost(context.Background(), &pb.PostId{Id: id})

	if err != nil {
		log.Fatalf("Error happened while deleting: %v\n", err)
	}

	log.Println("Post was deleted")
}
