package routes

import (
	"starter/httpd/handlers"

	"github.com/gin-gonic/gin"
)

func todoRoute(route *gin.Engine) {
	route.POST("/api/get-todo", handlers.GetData)
}
