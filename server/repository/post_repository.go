package repository

import (
	"context"
	"fmt"
	pb "github.com/neandermenezes/gRPC-Upvote-Service/proto/pb"
	"github.com/neandermenezes/gRPC-Upvote-Service/server/entity"
	"go.mongodb.org/mongo-driver/bson"
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
	ReadPost(id primitive.ObjectID, ctx context.Context) (*pb.Post, error)
	UpdatePost(id primitive.ObjectID, data *entity.PostItem, ctx context.Context) (*emptypb.Empty, error)
	DeletePost(id primitive.ObjectID, ctx context.Context) (*emptypb.Empty, error)
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

func (r *repo) ReadPost(id primitive.ObjectID, ctx context.Context) (*pb.Post, error) {
	data := &entity.PostItem{}
	filter := bson.M{"_id": id}

	res := Collection.FindOne(ctx, filter)

	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find the blog with id provided",
		)
	}

	return entity.DocumentToBlog(data), nil
}

func (r *repo) UpdatePost(id primitive.ObjectID, data *entity.PostItem, ctx context.Context) (*emptypb.Empty, error) {
	res, err := Collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": data},
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Could not update because %v\n", err),
		)
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find blog with Id",
		)
	}

	return &emptypb.Empty{}, nil
}

func (r *repo) DeletePost(id primitive.ObjectID, ctx context.Context) (*emptypb.Empty, error) {
	res, err := Collection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete object in MongoDB: %v", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Blog was not found",
		)
	}

	return &emptypb.Empty{}, nil
}

func (r *repo) UpvotePost(id primitive.ObjectID, ctx context.Context) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
