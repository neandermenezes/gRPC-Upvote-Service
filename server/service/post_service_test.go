package service

import (
	"context"
	pb "github.com/neandermenezes/gRPC-Upvote-Service/proto/pb"
	"github.com/neandermenezes/gRPC-Upvote-Service/server/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
)

type MockRepository struct {
	mock.Mock
}

var mockRepo = new(MockRepository)

func (m *MockRepository) CreatePost(data *entity.PostItem, ctx context.Context) (*pb.PostId, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*pb.PostId), args.Error(1)
}

func (m *MockRepository) ReadPost(id primitive.ObjectID, ctx context.Context) (*pb.Post, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*pb.Post), args.Error(1)
}

func (m *MockRepository) UpdatePost(id primitive.ObjectID, data *entity.PostItem, ctx context.Context) (*emptypb.Empty, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*emptypb.Empty), args.Error(1)
}

func (m *MockRepository) DeletePost(id primitive.ObjectID, ctx context.Context) (*emptypb.Empty, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*emptypb.Empty), args.Error(1)
}

func (m *MockRepository) ListPosts() (*mongo.Cursor, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*mongo.Cursor), args.Error(1)
}

func (m *MockRepository) UpvotePost(id primitive.ObjectID, ctx context.Context) (*emptypb.Empty, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*emptypb.Empty), args.Error(1)
}

func TestService_CreatePost(t *testing.T) {
	id := primitive.NewObjectID().Hex()
	mockRepo.On("CreatePost").Return(&pb.PostId{Id: id}, nil)

	testService := NewPostService(mockRepo)

	newPost := &pb.Post{
		AuthorName: "test",
		Title:      "test",
		Content:    "test",
	}

	res, err := testService.CreatePost(newPost, context.Background())
	if err != nil {
		t.Errorf("Could not create post: %v\n", err)
	}

	mockRepo.AssertExpectations(t)
	assert.Equal(t, id, res.Id)
}

func TestService_ReadPost(t *testing.T) {
	id := &pb.PostId{Id: primitive.NewObjectID().Hex()}

	expected := &pb.Post{
		Id:         id.Id,
		AuthorName: "teste",
		Title:      "teste",
		Content:    "teste",
		LikeCount:  0,
	}

	mockRepo.On("ListPosts").Return(expected, nil)
	testService := NewPostService(mockRepo)

	res, err := testService.ReadPost(id, context.Background())
	if err != nil {
		t.Errorf("Could not read post: %v\n", err)
	}

	mockRepo.AssertExpectations(t)
	assert.Equal(t, id.Id, res.Id)
	assert.Equal(t, expected.AuthorName, res.AuthorName)
	assert.Equal(t, expected.Title, res.Title)
	assert.Equal(t, expected.Content, res.Content)
}

func TestService_UpdatePost(t *testing.T) {
	id := &pb.PostId{Id: primitive.NewObjectID().Hex()}
	post := &pb.Post{
		Id:      id.Id,
		Title:   "teste",
		Content: "teste",
	}

	mockRepo.On("UpdatePost").Return(&emptypb.Empty{}, nil)
	testService := NewPostService(mockRepo)

	res, err := testService.UpdatePost(post, context.Background())

	if err != nil {
		t.Errorf("Could not update post: %v\n", err)
	}

	mockRepo.AssertExpectations(t)
	assert.Equal(t, &emptypb.Empty{}, res)
}

func TestService_DeletePost(t *testing.T) {
	id := &pb.PostId{Id: primitive.NewObjectID().Hex()}

	mockRepo.On("DeletePost").Return(&emptypb.Empty{}, nil)
	testService := NewPostService(mockRepo)

	res, err := testService.DeletePost(id, context.Background())

	if err != nil {
		t.Errorf("Could not delete post: %v\n", err)
	}

	mockRepo.AssertExpectations(t)
	assert.Equal(t, &emptypb.Empty{}, res)
}

//func TestService_ListPosts(t *testing.T) {
//	mockRepo.On("ListPosts").Return(&mongo.Cursor{}, nil)
//	testService := NewPostService(mockRepo)
//
//	err := testService.ListPosts(nil, nil)
//
//	if err != nil {
//		t.Errorf("Could not list posts: %v\n", err)
//	}
//
//	mockRepo.AssertExpectations(t)
//}

func TestService_UpvotePost(t *testing.T) {
	id := &pb.PostId{Id: primitive.NewObjectID().Hex()}

	mockRepo.On("UpvotePost").Return(&emptypb.Empty{}, nil)
	testService := NewPostService(mockRepo)

	res, err := testService.UpvotePost(id, context.Background())

	if err != nil {
		t.Errorf("Could not delete post: %v\n", err)
	}

	mockRepo.AssertExpectations(t)
	assert.Equal(t, &emptypb.Empty{}, res)
}
