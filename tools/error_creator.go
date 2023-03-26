package tools

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// CreateError позволяет вернуть кастомную ошибку
func CreateError(code int, requestErr error, c *gin.Context) {
	c.JSON(code, gin.H{
		"code":  code,
		"error": requestErr.Error(),
	})
	err := c.AbortWithError(code, errors.New(requestErr.Error()))

	if err != nil {
		println(err)
	}
}
