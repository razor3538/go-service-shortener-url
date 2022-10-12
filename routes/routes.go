package routes

import (
	"example.com/m/v2/internal/app/api"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	shortenAPI := api.NewShortURLAPI()

	r.POST("/short-url", shortenAPI.ShortenURL)
	r.GET("/short-url/:id", shortenAPI.GetFullURL)

	return r
}
