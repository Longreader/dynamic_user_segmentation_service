package handlers

import (
	_ "github.com/Longreader/dynamic_user_segmentation_service.git/docs"
	"github.com/Longreader/dynamic_user_segmentation_service.git/service"
	"github.com/gin-gonic/gin"

	// _ "github.com/swaggo/echo-swagger/example/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: *services}
}

func (h *Handler) InitRouter() *gin.Engine {
	// Инициализация роутера приложения
	// Подключение API Endpoints
	r := gin.New()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := r.Group("/api/v1")
	{
		segment := api.Group("/segments")
		{
			segment.POST("/", h.postSegment)
			segment.GET("/:segment", h.getSegment)
			segment.DELETE("/:segment", h.deleteSegment)
		}
		user := api.Group("/users")
		{
			user.POST("/", h.postUser)
			user.GET("/:id", h.getUser)
			user.DELETE("/:id", h.deleteUser)
			user.POST("/add", h.postComarison)
			user.GET("/active/:id", h.getActive)
		}

	}

	return r

}
