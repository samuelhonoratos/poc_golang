package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

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
	api := s.router.Group("/api")
	{
		ps := services.NewPostService(s.db)
		pc := controllers.NewPostController(ps)

		api.GET("/posts", pc.Index)
		api.POST("/post", pc.Create)
	}
}
