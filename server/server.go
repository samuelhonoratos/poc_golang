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
	api := s.router.Group("/api")
	{
		post_controller := post_controller.New(s.db)

		api.POST("/users", post_controller.Create)
	}

	s.router.Run(":8080")
}
