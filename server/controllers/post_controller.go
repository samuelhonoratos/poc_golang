package post_controller

import (
	"net/http"
	"orcamento/server/dto"

	post_service "orcamento/server/services"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	service post_service.PostService
}

func New(service post_service.PostService) *PostController {
	return &PostController{service: service}
}

func (pc *PostController) Index(context *gin.Context) {
	posts, err := pc.service.Index()

	if err != nil {
		context.AbortWithStatusJSON(http.StatusCreated, dto.ErrorResponseDTO{
			Message: err.Error(),
		})

		return
	}

	context.AbortWithStatusJSON(http.StatusCreated, gin.H{"posts": posts})
}

func (pc *PostController) Create(context *gin.Context) {
	postCreateDTO := dto.CreatePostDTO{}
	context.BindJSON(&postCreateDTO)

	post, err := pc.service.Create(postCreateDTO)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusCreated, dto.ErrorResponseDTO{
			Message: err.Error(),
		})

		return
	}

	context.AbortWithStatusJSON(http.StatusCreated, post)
}
