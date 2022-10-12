package api

import (
	"example.com/m/v2/internal/app/models"
	"example.com/m/v2/services"
	"example.com/m/v2/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ShortUrlAPI struct{}

func NewShortUrlAPI() *ShortUrlAPI {
	return &ShortUrlAPI{}
}

var urlService = services.NewUrlService()

func (sua *ShortUrlAPI) ShortenUrl(c *gin.Context) {
	var body models.ShortenUrl

	if err := tools.RequestBinderBody(&body, c); err != nil {
		return
	}

	urlModel, err := urlService.Save(body)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusCreated, urlModel.ShortUrl)
}

func (sua *ShortUrlAPI) GetFullUrl(c *gin.Context) {
	var path models.PathID

	if err := tools.RequestBinderURI(&path, c); err != nil {
		return
	}

	urlModel, err := urlService.Get(path.ID)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, urlModel.FullUrl)
}
