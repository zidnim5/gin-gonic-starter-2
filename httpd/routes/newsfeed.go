package routes

import (
	"starter/httpd/handlers"

	"github.com/gin-gonic/gin"
)

func newsfeedRoute(route *gin.Engine) {
	route.POST("/api/get-newsfeed", handlers.GetNewsFeed)
}
