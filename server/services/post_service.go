package services

import (
	"orcamento/domain/entity"
	"orcamento/repository"
	"orcamento/server/dto"
)

type PostService struct {
	repository *repository.PostRepository
}

func NewPostService(repository *repository.PostRepository) *PostService {
	return &PostService{repository: repository}
}

func (ps *PostService) Index() ([]entity.Post, error) {
	var posts []entity.Post

	posts, err := ps.repository.GetAllPosts()

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (ps *PostService) Create(payload dto.CreatePostDTO) (entity.Post, error) {
	post := entity.Post{
		Title: payload.Title,
		Body:  payload.Body,
	}

	err := ps.repository.Create(&post)

	if err != nil {
		return entity.Post{}, err
	}

	return post, nil
}
