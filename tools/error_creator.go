package tools

import (
	"errors"
	"github.com/gin-gonic/gin"
)

// CreateError позволяет вернуть кастомную ошибку
func CreateError(code int, err error, c *gin.Context) {
	c.JSON(code, gin.H{
		"code":  code,
		"error": err.Error(),
	})
	err = c.AbortWithError(code, errors.New(err.Error()))
	
	if err != nil {
		println(err)
	}
}
