package repository

import (
	"orcamento/domain/entity"

	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (pr *PostRepository) GetAllPosts() ([]entity.Post, error) {
	var posts []entity.Post

	res := pr.db.Find(&posts)

	if res.Error != nil {
		return nil, res.Error
	}

	return posts, nil
}

func (pr *PostRepository) Create(post *entity.Post) error {
	res := pr.db.Create(post)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
