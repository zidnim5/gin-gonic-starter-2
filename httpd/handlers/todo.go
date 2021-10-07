package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "list of data here",
	})
}
