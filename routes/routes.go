package routes

import (
	"example.com/m/v2/internal/app/api"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	shortenAPI := api.NewShortUrlAPI()

	r.POST("/short-url", shortenAPI.ShortenUrl)
	r.GET("/short-url/:id", shortenAPI.GetFullUrl)

	return r
}
