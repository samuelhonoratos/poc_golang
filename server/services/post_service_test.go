package services

import (
	"errors"
	"orcamento/domain/entity"
	"orcamento/server/dto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPostRepository is a mock type for the PostRepository type
type MockPostRepository struct {
	mock.Mock
}

func (m *MockPostRepository) GetAllPosts() ([]entity.Post, error) {
	args := m.Called()

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]entity.Post), args.Error(1)
}

func (m *MockPostRepository) Create(post *entity.Post) error {
	args := m.Called(post)
	return args.Error(0)
}

func TestPostService_Index(t *testing.T) {
	mockRepo := new(MockPostRepository)
	service := NewPostService(mockRepo)

	mockPosts := []entity.Post{
		{Title: "Post 1", Body: "Body 1"},
		{Title: "Post 2", Body: "Body 2"},
	}

	mockRepo.On("GetAllPosts").Return(mockPosts, nil)

	posts, err := service.Index()

	assert.NoError(t, err)
	assert.Equal(t, mockPosts, posts)
	mockRepo.AssertExpectations(t)
}

func TestPostService_Index_Error(t *testing.T) {
	mockRepo := new(MockPostRepository)
	service := NewPostService(mockRepo)

	mockPosts := []entity.Post{}
	mockRepo.On("GetAllPosts").Return(mockPosts, errors.New("some error"))

	posts, err := service.Index()

	assert.Error(t, err)
	assert.Nil(t, posts)
	mockRepo.AssertExpectations(t)
}

func TestPostService_Create(t *testing.T) {
	mockRepo := new(MockPostRepository)
	service := NewPostService(mockRepo)

	payload := dto.CreatePostDTO{
		Title: "New Post",
		Body:  "New Body",
	}

	expectedPost := entity.Post{
		Title: payload.Title,
		Body:  payload.Body,
	}

	mockRepo.On("Create", &expectedPost).Return(nil)

	post, err := service.Create(payload)

	assert.NoError(t, err)
	assert.Equal(t, expectedPost, post)
	mockRepo.AssertExpectations(t)
}

func TestPostService_Create_Error(t *testing.T) {
	mockRepo := new(MockPostRepository)
	service := NewPostService(mockRepo)

	payload := dto.CreatePostDTO{
		Title: "New Post",
		Body:  "New Body",
	}

	expectedPost := entity.Post{
		Title: payload.Title,
		Body:  payload.Body,
	}

	mockRepo.On("Create", &expectedPost).Return(errors.New("some error"))

	post, err := service.Create(payload)

	assert.Error(t, err)
	assert.Equal(t, entity.Post{}, post)
	mockRepo.AssertExpectations(t)
}
