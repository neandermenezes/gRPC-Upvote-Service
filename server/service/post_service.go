package service

import (
	"context"
	pb "github.com/neandermenezes/gRPC-Upvote-Service/proto/pb"
	"github.com/neandermenezes/gRPC-Upvote-Service/server/entity"
	"github.com/neandermenezes/gRPC-Upvote-Service/server/repository"
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
	//TODO implement me
	panic("implement me")
}

func (s *service) UpdatePost(in *pb.Post, ctx context.Context) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeletePost(in *pb.PostId, ctx context.Context) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) ListPosts(in *emptypb.Empty, stream pb.PostService_ListPostsServer) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) UpvotePost(in *pb.PostId, ctx context.Context) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}
