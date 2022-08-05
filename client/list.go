package main

import (
	"context"
	pb "github.com/neandermenezes/gRPC-Upvote-Service/proto/pb"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
)

func listPosts(c pb.PostServiceClient) {
	log.Println("---listBlog was invoked---")

	stream, err := c.ListPosts(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while calling ListPosts: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something went wrong: %v\n", err)
		}

		log.Println(res)
	}
}
