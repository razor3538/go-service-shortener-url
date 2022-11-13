package api

import (
	"encoding/json"
	"errors"
	"example.com/m/v2/config"
	"fmt"
	"github.com/gin-gonic/gin"
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
	var userID string
	var byteString string

	var headerToken = c.GetHeader("Authorization")
	if headerToken == "" {
		var hash, err = tools.HashCookie()
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}

		byteString = fmt.Sprintf("%x", hash)

		userID = byteString

		c.Writer.Header().Set("Authorization", userID)

	} else {
		userID = headerToken

		c.Writer.Header().Set("Authorization", userID)
	}

	b, err := io.ReadAll(reader)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	urlString := string(b)

	urlModel, err := urlService.Save(urlString, userID)

	if err != nil && urlModel.FullURL != "" {
		c.Writer.WriteHeader(http.StatusConflict)

		c.Writer.Write([]byte(urlModel.FullURL))
		return
	}

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

	if err != nil && urlModel.FullURL != "" {
		c.JSON(http.StatusConflict, gin.H{
			"result": urlModel.ShortURL,
		})
		return
	}

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
	headerToken := c.GetHeader("Authorization")

	if headerToken == "" {
		tools.CreateError(http.StatusNoContent, errors.New("пустой токен"), c)
		return
	}
	userID := headerToken

	urlModel, err := urlService.GetByUserID(userID)

	if err != nil {
		tools.CreateError(http.StatusNoContent, err, c)
		return
	}

	c.JSON(http.StatusOK, urlModel)
}

func (sua *ShortURLAPI) SaveMany(c *gin.Context) {
	var body []models.SaveBatchURLRequest

	if err := tools.RequestBinderBody(&body, c); err != nil {
		return
	}

	urlModel, err := urlService.SaveMany(body)

	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusCreated, urlModel)
}

func (sua *ShortURLAPI) Ping(c *gin.Context) {
	if config.Env.BdConnection != "" {
		sqlDB, err := config.DB.DB()
		if err != nil {
			tools.CreateError(http.StatusInternalServerError, err, c)
			return
		}
		if err = sqlDB.Ping(); err != nil {
			err := sqlDB.Close()
			if err != nil {
				tools.CreateError(http.StatusInternalServerError, err, c)
				return
			}
		}
		c.Writer.WriteHeader(http.StatusOK)
	}
}
