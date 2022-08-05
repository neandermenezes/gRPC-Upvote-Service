package service

import (
	pb "github.com/neandermenezes/gRPC-Upvote-Service/proto/pb"
	"github.com/neandermenezes/gRPC-Upvote-Service/server/entity"
	"github.com/neandermenezes/gRPC-Upvote-Service/server/repository"
	"log"
)

type PostService interface {
	CreatePost(in *pb.Post) (*pb.PostId, error)
}

type service struct{}

var postRepository repository.PostRepository

func NewPostService(repository repository.PostRepository) PostService {
	postRepository = repository
	return &service{}
}

func (*service) CreatePost(in *pb.Post) (*pb.PostId, error) {
	log.Println("CreatePost service was invoked")

	data := entity.PostItem{
		AuthorName: in.AuthorName,
		Title:      in.Title,
		Content:    in.Content,
		LikeCount:  0,
	}

	res, err := postRepository.CreatePost(&data)

	if err != nil {
		return nil, err
	}

	return res, err
}
