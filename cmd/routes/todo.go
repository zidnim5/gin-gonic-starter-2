package routes

import (
	"starter/cmd/handlers"
	"github.com/gin-gonic/gin"
)

/**
 * TODO: routes of todo 
 */
func TodoRoute(route *gin.Engine) {
	route.GET("/api/todo/:id", handlers.FindTodo)
	route.POST("/api/todo", handlers.StoreTodo)	
}
