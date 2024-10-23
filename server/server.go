package server

import (
	post_controller "orcamento/server/controllers"
	post_service "orcamento/server/services"

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
		ps := post_service.New(s.db)
		post_controller := post_controller.New(*ps)

		api.GET("/posts", post_controller.Index)
		api.POST("/post", post_controller.Create)
	}
}
