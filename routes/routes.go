package routes

import (
	"example.com/m/v2/routes/middleware"

	"github.com/gin-gonic/gin"

	"example.com/m/v2/internal/app/api"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	shortenAPI := api.NewShortURLAPI()

	r.Use(middleware.GzipMiddleware)

	r.POST("/", shortenAPI.ShortenURL)

	r.POST("/api/shorten", shortenAPI.ReturnFullURL)

	r.GET("/:id", shortenAPI.GetFullURL)

	return r
}
