package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"orcamento/domain/entity"
	"orcamento/server/controllers"
	"orcamento/server/services"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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

func TestIndex(t *testing.T) {
	t.Run("it should return 200", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		r := gin.Default()

		mockRepo := new(MockPostRepository)
		ps := services.NewPostService(mockRepo)
		pc := controllers.NewPostController(ps)

		mockRepo.On("GetAllPosts").Return([]entity.Post{{ID: 1, Title: "Post 1"}}, nil)

		r.GET("/api/posts", pc.Index)

		req, _ := http.NewRequest(http.MethodGet, "/api/posts", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		expected := `{"posts":[{"id":1,"title":"Post 1","body":""}]}`
		assert.JSONEq(t, expected, w.Body.String())

		mockRepo.AssertExpectations(t)
	})
}
