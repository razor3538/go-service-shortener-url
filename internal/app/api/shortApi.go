package api

import (
	"example.com/m/v2/services"
	"example.com/m/v2/tools"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type ShortURLAPI struct{}

func NewShortURLAPI() *ShortURLAPI {
	return &ShortURLAPI{}
}

var urlService = services.NewURLService()

func (sua *ShortURLAPI) ShortenURL(c *gin.Context) {
	var body string

	//err := json.NewDecoder(c.Request.Body).Decode(&body)
	//if err != nil {
	//	println("error")
	//	tools.CreateError(http.StatusBadRequest, err, c)
	//	return
	//}

	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	myString := string(b)

	println(myString)
	println(myString)
	println(myString)
	println("myString")

	urlModel, err := urlService.Save(body)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}
	c.Writer.WriteHeader(http.StatusCreated)
	c.Writer.Write([]byte(urlModel.ShortURL))

	println(urlModel.FullURL)
	println(urlModel.ShortURL)
}

func (sua *ShortURLAPI) GetFullURL(c *gin.Context) {
	name := c.Param("id")

	//_, err := urlService.Get(name)
	urlModel, err := urlService.Get(name)
	println(urlModel.ID)
	println(urlModel.FullURL)
	println(urlModel.ShortURL)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.Writer.Header().Set("Location", urlModel.FullURL)

	c.JSON(http.StatusTemporaryRedirect, nil)
}
