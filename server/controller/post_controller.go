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

func ReadBlog(ctx context.Context, in *pb.PostId) (*pb.Post, error) {
	return nil, nil
}

func UpdateBlog(ctx context.Context, in *pb.Post) (*emptypb.Empty, error) {
	return nil, nil
}

func DeleteBlog(ctx context.Context, in *pb.PostId) (*emptypb.Empty, error) {
	return nil, nil
}

func ListBlogs(in *emptypb.Empty, stream pb.PostService_ListBlogsServer) error {
	return nil
}

func UpvotePost(ctx context.Context, in *pb.PostId) (*emptypb.Empty, error) {
	return nil, nil
}
