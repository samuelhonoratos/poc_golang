package repository

import (
	"orcamento/domain/entity"

	"gorm.io/gorm"
)

type PostRepository interface {
	GetAllPosts() ([]entity.Post, error)
	Create(post *entity.Post) error
}

type PostRepositoryImpl struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepositoryImpl {
	return &PostRepositoryImpl{db: db}
}

func (pr *PostRepositoryImpl) GetAllPosts() ([]entity.Post, error) {
	var posts []entity.Post

	res := pr.db.Find(&posts)

	if res.Error != nil {
		return nil, res.Error
	}

	return posts, nil
}

func (pr *PostRepositoryImpl) Create(post *entity.Post) error {
	res := pr.db.Create(post)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
