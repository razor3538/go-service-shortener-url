package api

import (
	"example.com/m/v2/internal/app/models"
	"example.com/m/v2/services"
	"example.com/m/v2/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ShortURLAPI struct{}

func NewShortURLAPI() *ShortURLAPI {
	return &ShortURLAPI{}
}

var urlService = services.NewURLService()

func (sua *ShortURLAPI) ShortenURL(c *gin.Context) {
	var body models.ShortenURL

	if err := tools.RequestBinderBody(&body, c); err != nil {
		return
	}

	urlModel, err := urlService.Save(body)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}
	c.Writer.WriteHeader(http.StatusCreated)

	c.Writer.Write([]byte(urlModel.ShortURL))
}

func (sua *ShortURLAPI) GetFullURL(c *gin.Context) {
	name := c.Param("id")

	_, err := urlService.Get(name)
	//urlModel, err := urlService.Get(name)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}
	println("c.GetHeader(Location)")
	println(c.GetHeader("Location"))

	c.Writer.Header().Set("Location", "https://google.com")
	println("c.GetHeader(Location)")

	//c.Writer.Header().Add("Location", urlModel.FullURL)

	c.JSON(http.StatusTemporaryRedirect, nil)
}
