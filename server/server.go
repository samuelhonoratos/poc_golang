package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"orcamento/repository"
	"orcamento/server/controllers"
	"orcamento/server/services"
)

type Server struct {
	router *gin.Engine
	db     *gorm.DB
}

func New(db *gorm.DB) *Server {
	return &Server{
		router: gin.Default(),
		db:     db,
	}
}

func (s *Server) Routes() {
	s.ApiGroupRoute()

	s.router.Run(":8080")
}

func (s *Server) ApiGroupRoute() {
	pr := repository.NewPostRepository(s.db)
	ps := services.NewPostService(pr)
	pc := controllers.NewPostController(ps)

	api := s.router.Group("/api")
	{
		api.GET("/posts", pc.Index)
		api.POST("/post", pc.Create)
	}
}
