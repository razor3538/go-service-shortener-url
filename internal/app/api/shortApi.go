package api

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"example.com/m/v2/internal/app/models"
	"example.com/m/v2/services"
	"example.com/m/v2/tools"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

type ShortURLAPI struct{}

func NewShortURLAPI() *ShortURLAPI {
	return &ShortURLAPI{}
}

var urlService = services.NewURLService()

func (sua *ShortURLAPI) ShortenURL(c *gin.Context) {
	var urlString string

	if strings.Contains(c.Request.Header.Get("Accept-Encoding"), "gzip") {
		gz, err := gzip.NewReader(c.Request.Body)
		if err != nil {
			tools.CreateError(http.StatusBadRequest, errors.New("ошибка номер 1"), c)
			return
		}
		defer gz.Close()

		b, err := io.ReadAll(gz)
		if err != nil {
			tools.CreateError(http.StatusBadRequest, errors.New("ошибка номер 2"), c)
			return
		}
		urlString = string(b)

		tools.CreateError(http.StatusBadRequest, errors.New(urlString), c)
		return

	}

	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, errors.New("ошибка номер 3"), c)
		return
	}
	urlString = string(b)

	urlModel, err := urlService.Save(urlString)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}
	c.Writer.WriteHeader(http.StatusCreated)
	c.Writer.Write([]byte(urlModel.ShortURL))
}

func (sua *ShortURLAPI) ReturnFullURL(c *gin.Context) {
	var body models.URLRequestModel

	if err := tools.RequestBinderBody(&body, c); err != nil {
		return
	}

	urlModel, err := urlService.GetByFullURL(body.URL)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	jsonModel, err := json.Marshal(gin.H{
		"result": urlModel.ShortURL,
	})

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	println(jsonModel)

	c.JSON(http.StatusCreated, gin.H{
		"result": urlModel.ShortURL,
	})
}

func (sua *ShortURLAPI) GetFullURL(c *gin.Context) {
	name := c.Param("id")

	//_, err := urlService.Get(name)
	urlModel, err := urlService.Get(name)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.Writer.Header().Set("Location", urlModel.FullURL)

	c.JSON(http.StatusTemporaryRedirect, nil)
}
