package post_service

import (
	"orcamento/domain/entity"

	"gorm.io/gorm"
)

type PostService struct {
	db *gorm.DB
}

func New(db *gorm.DB) *PostService {
	return &PostService{db: db}
}

func (ps *PostService) Index() ([]entity.Post, string) {
	var posts []entity.Post

	res := ps.db.Find(&posts)

	if res.Error != nil {
		return nil, res.Error.Error()
	}

	return posts, ""
}
