package main

import (
	"context"
	pb "github.com/neandermenezes/gRPC-Upvote-Service/proto/pb"
	"log"
)

func readPost(c pb.PostServiceClient, id string) *pb.Post {
	log.Println("---Read post was invoked---")

	req := &pb.PostId{
		Id: id,
	}
	res, err := c.ReadPost(context.Background(), req)

	if err != nil {
		log.Printf("Error happened while reading: %v\n", err)
	}

	log.Printf("Post was read: %v\n", res)
	return res
}
