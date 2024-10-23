package server

import (
	post_controller "orcamento/server/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		post_controller := post_controller.New(s.db)

		api.GET("/posts", post_controller.Index)
		api.POST("/post", post_controller.Create)
	}
}
