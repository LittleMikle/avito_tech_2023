package handler

import (
	"github.com/LittleMikle/avito_tech_2023/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		segment := api.Group("/segments")
		{
			segment.POST("/create", h.createSegment)
			segment.DELETE("/:id", h.deleteSegment)

		}

		usersSeg := api.Group("/users")
		{
			usersSeg.POST("/create", h.createUsersSeg)
			usersSeg.GET("/:id", h.getUsersSeg)
			usersSeg.GET("/history/:id", h.getHistory)
			usersSeg.DELETE("/schedule/:id", h.scheduleDelete)
			usersSeg.POST("/random", h.randomCreate)
		}

		return router
	}
}
