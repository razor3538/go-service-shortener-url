package middleware

import (
	"compress/gzip"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GzipMiddleware проверяет были ли данные в запросе сжаты и если да, то распаковывает их
func GzipMiddleware(c *gin.Context) {
	if c.Request.Header.Get(`Content-Encoding`) == `gzip` {
		gz, err := gzip.NewReader(c.Request.Body)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}
		c.Request.Body = gz
		defer gz.Close()
	}
	c.Next()
}
