package handlers

import (
	"net/http"
	"strings"

	_ "github.com/Longreader/dynamic_user_segmentation_service.git/docs"
	"github.com/Longreader/dynamic_user_segmentation_service.git/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: *services}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := viper.GetString("token")

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" || parts[1] != token {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Next()
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	// Инициализация роутера приложения
	// Подключение API Endpoints

	// gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := r.Group("/api/v1", AuthMiddleware())
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
		utils := api.Group("/utils")
		{
			utils.GET("/audit/:date", h.dowloadAudit)
		}

	}

	return r

}
