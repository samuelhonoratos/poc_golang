package services

import (
	"orcamento/domain/entity"
	"orcamento/server/dto"

	"gorm.io/gorm"
)

type PostService struct {
	db *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{db: db}
}

func (ps *PostService) Index() ([]entity.Post, error) {
	var posts []entity.Post

	res := ps.db.Find(&posts)

	if res.Error != nil {
		return nil, res.Error
	}

	return posts, nil
}

func (ps *PostService) Create(payload dto.CreatePostDTO) (entity.Post, error) {
	post := entity.Post{
		Title: payload.Title,
		Body:  payload.Body,
	}

	res := ps.db.Create(&post)

	if res.Error != nil {
		return entity.Post{}, res.Error
	}

	return post, nil
}
