package controller

import (
	"context"
	pb "github.com/neandermenezes/gRPC-Upvote-Service/proto/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PostServer struct {
	pb.PostServiceServer
}

func (s *PostServer) CreateBlog(ctx context.Context, in *pb.Post) (*pb.PostId, error) {
	return nil, nil
}

func (s *PostServer) ReadBlog(ctx context.Context, in *pb.PostId) (*pb.Post, error) {
	return nil, nil
}

func (s *PostServer) UpdateBlog(ctx context.Context, in *pb.Post) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *PostServer) DeleteBlog(ctx context.Context, in *pb.PostId) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *PostServer) ListBlogs(in *emptypb.Empty, stream pb.PostService_ListBlogsServer) error {
	return nil
}

func (s *PostServer) UpvotePost(ctx context.Context, in *pb.PostId) (*emptypb.Empty, error) {
	return nil, nil
}
