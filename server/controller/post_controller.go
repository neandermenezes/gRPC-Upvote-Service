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
	res, err := postService.CreatePost(in)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *PostServer) ReadPost(ctx context.Context, in *pb.PostId) (*pb.Post, error) {
	return nil, nil
}

func (s *PostServer) UpdatePost(ctx context.Context, in *pb.Post) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *PostServer) DeletePost(ctx context.Context, in *pb.PostId) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *PostServer) ListPosts(in *emptypb.Empty, stream pb.PostService_ListPostsServer) error {
	return nil
}

func (s *PostServer) UpvotePost(ctx context.Context, in *pb.PostId) (*emptypb.Empty, error) {
	return nil, nil
}
