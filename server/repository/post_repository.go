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
	"google.golang.org/protobuf/types/known/emptypb"
)

type repo struct{}

var (
	Collection *mongo.Collection
)

type PostRepository interface {
	CreatePost(data *entity.PostItem, ctx context.Context) (*pb.PostId, error)
	ReadPost(id primitive.ObjectID, ctx context.Context) (*entity.PostItem, error)
	UpdatePost(data *entity.PostItem, ctx context.Context) (*emptypb.Empty, error)
	DeletePost(id *pb.PostId, ctx context.Context) (*emptypb.Empty, error)
	//ListPosts(in *emptypb.Empty, stream pb.PostService_ListPostsServer) error
	UpvotePost(id primitive.ObjectID, ctx context.Context) (*emptypb.Empty, error)
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

func (r *repo) ReadPost(id primitive.ObjectID, ctx context.Context) (*entity.PostItem, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repo) UpdatePost(data *entity.PostItem, ctx context.Context) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repo) DeletePost(id *pb.PostId, ctx context.Context) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (r *repo) UpvotePost(id primitive.ObjectID, ctx context.Context) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
