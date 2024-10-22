package post_controller

import (
	"orcamento/domain/entity"
	"orcamento/server/dto"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostController struct {
	db *gorm.DB
}

func New(db *gorm.DB) *PostController {
	return &PostController{db: db}
}

func (pc *PostController) Create(context *gin.Context) {
	postCreateDTO := dto.CreatePostDTO{}
	context.BindJSON(&postCreateDTO)

	post := entity.Post{
		Title: postCreateDTO.Title,
		Body:  postCreateDTO.Body,
	}

	pc.db.Create(&post)

	context.JSON(201, gin.H{"message": "Usuário criado"})
}
