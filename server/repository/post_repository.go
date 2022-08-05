package repository

import (
	"context"
	"fmt"
	pb "github.com/neandermenezes/gRPC-Upvote-Service/proto/pb"
	"github.com/neandermenezes/gRPC-Upvote-Service/server/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type repo struct{}

var (
	Collection *mongo.Collection
)

type PostRepository interface {
	CreatePost(in *entity.PostItem, ctx context.Context) (*pb.PostId, error)
}

func NewPostRepository() PostRepository {
	return &repo{}
}

func (r *repo) CreatePost(data *entity.PostItem, ctx context.Context) (*pb.PostId, error) {
	res, err := Collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error REPOSITORY: %v\n", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot convert to OID",
		)
	}

	return &pb.PostId{Id: oid.Hex()}, nil
}
