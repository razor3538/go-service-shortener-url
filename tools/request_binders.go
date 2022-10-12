package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RequestTest interface {
	Body(model interface{}, c *gin.Context) error
}

func RequestBinderBody(model interface{}, c *gin.Context) error {
	if err := c.ShouldBind(model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return err
	}
	return nil
}

func RequestBinderURI(model interface{}, c *gin.Context) error {
	if err := c.ShouldBindUri(model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return err
	}
	return nil
}
