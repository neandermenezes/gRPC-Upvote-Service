package service

import (
	"context"
	"fmt"
	pb "github.com/neandermenezes/gRPC-Upvote-Service/proto/pb"
	"github.com/neandermenezes/gRPC-Upvote-Service/server/entity"
	"github.com/neandermenezes/gRPC-Upvote-Service/server/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type PostService interface {
	CreatePost(in *pb.Post, ctx context.Context) (*pb.PostId, error)
	ReadPost(in *pb.PostId, ctx context.Context) (*pb.Post, error)
	UpdatePost(in *pb.Post, ctx context.Context) (*emptypb.Empty, error)
	DeletePost(in *pb.PostId, ctx context.Context) (*emptypb.Empty, error)
	ListPosts(in *emptypb.Empty, stream pb.PostService_ListPostsServer) error
	UpvotePost(in *pb.PostId, ctx context.Context) (*emptypb.Empty, error)
}

type service struct{}

var postRepository repository.PostRepository

func NewPostService(repository repository.PostRepository) PostService {
	postRepository = repository
	return &service{}
}

func (*service) CreatePost(in *pb.Post, ctx context.Context) (*pb.PostId, error) {
	log.Println("CreatePost service was invoked")

	data := entity.PostItem{
		AuthorName: in.AuthorName,
		Title:      in.Title,
		Content:    in.Content,
		LikeCount:  0,
	}

	res, err := postRepository.CreatePost(&data, ctx)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *service) ReadPost(in *pb.PostId, ctx context.Context) (*pb.Post, error) {
	log.Println("ReadPost service was invoked")

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	res, err := postRepository.ReadPost(oid, ctx)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *service) UpdatePost(in *pb.Post, ctx context.Context) (*emptypb.Empty, error) {
	log.Println("UpdatePost service was invoked")

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	data := &entity.PostItem{
		Title:   in.Title,
		Content: in.Content,
	}

	res, err := postRepository.UpdatePost(oid, data, ctx)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *service) DeletePost(in *pb.PostId, ctx context.Context) (*emptypb.Empty, error) {
	log.Println("DeletePost service was invoked")

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	res, err := postRepository.DeletePost(oid, ctx)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *service) ListPosts(in *emptypb.Empty, stream pb.PostService_ListPostsServer) error {
	log.Println("ListPosts service was invoked")

	cursor, err := postRepository.ListPosts()
	ctx := context.Background()

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v\n", err),
		)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		data := &entity.PostItem{}
		err := cursor.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data from MongoDB: %v", err),
			)
		}

		stream.Send(entity.DocumentToBlog(data))
	}

	if err = cursor.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}

	return nil
}

func (s *service) UpvotePost(in *pb.PostId, ctx context.Context) (*emptypb.Empty, error) {
	log.Println("DeletePost service was invoked")

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	res, err := postRepository.UpvotePost(oid, ctx)

	if err != nil {
		return nil, err
	}

	return res, err
}
