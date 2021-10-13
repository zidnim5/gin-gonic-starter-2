package routes

import (
	"starter/cmd/handlers"
	"github.com/gin-gonic/gin"
)

func TodoRoute(route *gin.Engine) {
	route.GET("/api/todo", handlers.GetTodo)
	route.POST("/api/todo", handlers.StoreTodo)	
}
