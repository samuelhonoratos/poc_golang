package post_controller

import (
	"orcamento/domain/entity"

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
	post := entity.Post{Title: "Teste", Body: "Teste"}

	pc.db.Create(&post)

	context.JSON(200, "mensagem de retorno1")
}
