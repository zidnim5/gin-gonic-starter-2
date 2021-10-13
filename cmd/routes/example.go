package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func exampleRoute(route *gin.Engine) {
	route.POST("/api/get-example", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "oke",
		})
	})
}
