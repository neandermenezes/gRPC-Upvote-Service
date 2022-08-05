package main

import (
	pb "github.com/neandermenezes/gRPC-Upvote-Service/proto/pb"
	"github.com/neandermenezes/gRPC-Upvote-Service/server/controller"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	//ctx := context.Background()
	//
	//client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//err = client.Connect(ctx)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//collection = client.Database("upvote-db").Collection("posts")
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
