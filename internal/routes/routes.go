package routes

import (
	"github.com/gin-gonic/gin"
	"go-service-shortener-url/internal/app/api"
	"go-service-shortener-url/internal/routes/middleware"
)

// SetupRouter настраивает все роуты приложения, а так же устанавливает middleware
func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	shortenAPI := api.NewShortURLAPI()

	r.Use(middleware.GzipMiddleware)

	r.POST("/", shortenAPI.ShortenURL)

	r.DELETE("/api/user/urls", shortenAPI.DeleteURLs)

	r.POST("/api/shorten", shortenAPI.ReturnFullURL)

	r.POST("/api/shorten/batch", shortenAPI.SaveMany)

	r.GET("/:id", shortenAPI.GetFullURL)

	r.GET("/api/user/urls", shortenAPI.GetByUserID)

	r.GET("/ping", shortenAPI.Ping)

	return r
}
