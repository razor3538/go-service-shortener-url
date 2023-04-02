package api

import (
	"encoding/json"
	"errors"
	"example.com/m/v2/internal/app/models"
	config2 "example.com/m/v2/internal/config"
	"example.com/m/v2/internal/services"
	"example.com/m/v2/internal/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

// ShortURLAPI Структура обрабатываюващая обращения к API
type ShortURLAPI struct{}

// NewShortURLAPI возвращает указатель на структуру ShortURLAPI
// со всеми его методами
func NewShortURLAPI() *ShortURLAPI {
	return &ShortURLAPI{}
}

var urlService = services.NewURLService()

// DeleteURLs обработчик эндопоинта для удаления урлов по пользователю
func (sua *ShortURLAPI) DeleteURLs(c *gin.Context) {
	var reader = c.Request.Body

	b, err := io.ReadAll(reader)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	urlString := string(b)

	go urlService.Delete(tools.StringToSlice(urlString))

	c.Writer.WriteHeader(http.StatusAccepted)
}

// ShortenURL обработчик эндопоинта для сохранения урла полученного из строки
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

		_, errWrite := c.Writer.Write([]byte(urlModel.ShortURL))
		if errWrite != nil {
			CreateError(http.StatusBadRequest, err, c)
			return
		}
		return
	} else if err != nil {
		CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.Writer.WriteHeader(http.StatusCreated)

	_, err = c.Writer.Write([]byte(urlModel.ShortURL))

	if err != nil {
		return
	}
}

// ReturnFullURL сокращает урл полученный из JSON
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
		CreateError(http.StatusBadRequest, err, c)
		return
	}

	jsonModel, err := json.Marshal(gin.H{
		"result": urlModel.ShortURL,
	})

	if err != nil {
		CreateError(http.StatusBadRequest, err, c)
		return
	}

	println(jsonModel)

	c.JSON(http.StatusCreated, gin.H{
		"result": urlModel.ShortURL,
	})
}

// GetFullURL обработчик эндопоинта для сохранения урла полученного из JSON
func (sua *ShortURLAPI) GetFullURL(c *gin.Context) {
	name := c.Param("id")

	urlModel, err := urlService.Get(name)

	if err != nil {
		CreateError(http.StatusBadRequest, err, c)
		return
	}

	if urlModel.Deleted.Valid {
		c.Writer.WriteHeader(http.StatusGone)
	} else {
		c.Writer.Header().Set("Location", urlModel.FullURL)

		c.JSON(http.StatusTemporaryRedirect, nil)
	}
}

// GetByUserID обработчик эндопоинта для получения всех сохраненых пользователем урлов
func (sua *ShortURLAPI) GetByUserID(c *gin.Context) {
	headerToken := c.GetHeader("Authorization")

	if headerToken == "" {
		CreateError(http.StatusNoContent, errors.New("пустой токен"), c)
		return
	}
	userID := headerToken

	urlModel, err := urlService.GetByUserID(userID)

	if err != nil {
		CreateError(http.StatusNoContent, err, c)
		return
	}

	c.JSON(http.StatusOK, urlModel)
}

// SaveMany обработчик эндопоинта для сохранения множества урлов 1 запросом
func (sua *ShortURLAPI) SaveMany(c *gin.Context) {
	var body []models.SaveBatchURLRequest

	if err := tools.RequestBinderBody(&body, c); err != nil {
		return
	}

	urlModel, err := urlService.SaveMany(body)

	if err != nil {
		CreateError(http.StatusBadRequest, err, c)
		return
	}

	c.JSON(http.StatusCreated, urlModel)
}

// Ping обработчик эндопоинта для проверки работоспособности базы данных
func (sua *ShortURLAPI) Ping(c *gin.Context) {
	if config2.Env.BdConnection != "" {
		sqlDB, err := config2.DB.DB()
		if err != nil {
			CreateError(http.StatusInternalServerError, err, c)
			return
		}
		if err = sqlDB.Ping(); err != nil {
			err := sqlDB.Close()
			if err != nil {
				CreateError(http.StatusInternalServerError, err, c)
				return
			}
		}
		c.Writer.WriteHeader(http.StatusOK)
	}
}
