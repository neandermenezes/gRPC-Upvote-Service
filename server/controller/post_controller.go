package controller

import (
	"context"
	pb "github.com/neandermenezes/gRPC-Upvote-Service/proto/pb"
	"github.com/neandermenezes/gRPC-Upvote-Service/server/service"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PostServer struct {
	pb.PostServiceServer
}

var (
	postService service.PostService
)

func NewPostController(service service.PostService) PostServer {
	postService = service
	return PostServer{}
}

func (s *PostServer) CreatePost(ctx context.Context, in *pb.Post) (*pb.PostId, error) {
	res, err := postService.CreatePost(in, ctx)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PostServer) ReadPost(ctx context.Context, in *pb.PostId) (*pb.Post, error) {
	res, err := postService.ReadPost(in, ctx)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *PostServer) UpdatePost(ctx context.Context, in *pb.Post) (*emptypb.Empty, error) {
	res, err := postService.UpdatePost(in, ctx)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *PostServer) DeletePost(ctx context.Context, in *pb.PostId) (*emptypb.Empty, error) {
	res, err := postService.DeletePost(in, ctx)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *PostServer) ListPosts(in *emptypb.Empty, stream pb.PostService_ListPostsServer) error {
	err := postService.ListPosts(in, stream)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostServer) UpvotePost(ctx context.Context, in *pb.PostId) (*emptypb.Empty, error) {
	res, err := postService.UpvotePost(in, ctx)

	if err != nil {
		return nil, err
	}

	return res, err
}
