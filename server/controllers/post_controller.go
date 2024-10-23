package post_controller

import (
	"net/http"
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

func (pc *PostController) Index(context *gin.Context) {
	var posts []entity.Post

	res := pc.db.Find(&posts)

	if res.Error != nil {
		context.AbortWithStatusJSON(http.StatusCreated, dto.ErrorResponseDTO{
			Message: res.Error.Error(),
		})

		return
	}

	context.AbortWithStatusJSON(http.StatusCreated, gin.H{"posts": posts})
}

func (pc *PostController) Create(context *gin.Context) {
	postCreateDTO := dto.CreatePostDTO{}
	context.BindJSON(&postCreateDTO)

	post := entity.Post{
		Title: postCreateDTO.Title,
		Body:  postCreateDTO.Body,
	}

	res := pc.db.Create(&post)

	if res.Error != nil {
		context.AbortWithStatusJSON(http.StatusCreated, dto.ErrorResponseDTO{
			Message: res.Error.Error(),
		})

		return
	}

	context.AbortWithStatusJSON(http.StatusCreated, post)
}
