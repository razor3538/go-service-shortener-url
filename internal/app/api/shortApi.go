package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"

	"io"

	"net/http"

	"example.com/m/v2/internal/app/models"

	"example.com/m/v2/services"

	"example.com/m/v2/tools"
)

type ShortURLAPI struct{}

func NewShortURLAPI() *ShortURLAPI {
	return &ShortURLAPI{}
}

var urlService = services.NewURLService()

func (sua *ShortURLAPI) ShortenURL(c *gin.Context) {
	var reader = c.Request.Body
	var userId string
	var byteString string

	var _, err = c.Cookie("id")
	if err != nil {
		var hash, err = tools.HashCookie()
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}

		byteString = fmt.Sprintf("%x", hash)

		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "Authorization",
			Value:    byteString,
			Expires:  time.Now().Add(time.Hour * 24),
			HttpOnly: true,
			Secure:   false,
		})
		userId = ""

		c.Writer.Header().Set("Authorization", byteString)

	} else {
		cookie, err := c.Request.Cookie("id")
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}
		userId = cookie.Value

		c.Writer.Header().Set("Authorization", userId)
	}

	b, err := io.ReadAll(reader)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	urlString := string(b)

	urlModel, err := urlService.Save(urlString, userId)

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

	urlModel, err := urlService.Get(name)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.Writer.Header().Set("Location", urlModel.FullURL)

	c.JSON(http.StatusTemporaryRedirect, nil)
}

func (sua *ShortURLAPI) GetByUserID(c *gin.Context) {
	cookie, err := c.Request.Cookie("Authorization")
	println(c.GetHeader("Authorization"))
	println(c.GetHeader("Authorization"))
	println(c.GetHeader("Authorization"))
	println(c.GetHeader("Authorization"))

	if err != nil {
		println("dasdas")
		println(err.Error())
		println(err.Error())
		println(err.Error())
		tools.CreateError(http.StatusNoContent, err, c)
		return
	}
	userId := cookie.Value

	urlModel, err := urlService.GetByUserID(userId)

	if err != nil {
		tools.CreateError(http.StatusNoContent, err, c)
		return
	}

	c.JSON(http.StatusOK, urlModel)
}
