package routes

import (
	"example.com/m/v2/internal/app/api"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	shortenAPI := api.NewShortURLAPI()

	r.POST("/", shortenAPI.ShortenURL)
	r.GET("/:id", shortenAPI.GetFullURL)

	return r
}
