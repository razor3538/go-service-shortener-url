package tests

import (
	"example.com/m/v2/routes"
	"github.com/appleboy/gofight/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHomepageHandler(t *testing.T) {
	r := gofight.New()

	var shortURL string

	r.POST("/").
		SetBody("https://yandex.ru").
		Run(routes.SetupRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusCreated, r.Code)
			assert.Greater(t, r.Body.Len(), 0)
			shortURL = r.Body.String()
		})
	r.POST("/").
		SetBody("httpsы://google.com").
		Run(routes.SetupRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusBadRequest, r.Code)
		})
	r.POST("/").
		SetBody("httpsы://googlecom").
		Run(routes.SetupRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusBadRequest, r.Code)
		})
	r.POST("/").
		SetBody("google.com").
		Run(routes.SetupRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusBadRequest, r.Code)
		})

	r.GET(shortURL).
		Run(routes.SetupRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusTemporaryRedirect, r.Code)
		})

	r.GET("shortURL").
		Run(routes.SetupRouter(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusNotFound, r.Code)
		})
}
