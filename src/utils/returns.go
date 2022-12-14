package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CommonReturn(c *gin.Context, err interface{}, res interface{}) {

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"res": res,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"res": res,
	})
}
