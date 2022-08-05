package entity

import (
	pb "github.com/neandermenezes/gRPC-Upvote-Service/proto/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostItem struct {
	ID         primitive.ObjectID `bson:"id",omitempty`
	AuthorName string             `bson:"authorName"`
	Title      string             `bson:"title"`
	Content    string             `bson:"content"`
	LikeCount  uint32             `bson:"likeCount"`
}

func DocumentToBlog(data *PostItem) *pb.Post {
	return &pb.Post{
		Id:         data.ID.Hex(),
		AuthorName: data.AuthorName,
		Title:      data.Title,
		Content:    data.Content,
		LikeCount:  0,
	}
}
