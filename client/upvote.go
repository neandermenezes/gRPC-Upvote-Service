package main

import (
	"context"
	pb "github.com/neandermenezes/gRPC-Upvote-Service/proto/pb"
	"log"
)

func upvotePost(c pb.PostServiceClient, id string) {
	log.Println("---upvotePost was invoked---")

	req := &pb.PostId{Id: id}

	res, err := c.UpvotePost(context.Background(), req)

	if err != nil {
		log.Printf("Error happened while upvoting: %v\n", err)
	}

	log.Println(res)
}