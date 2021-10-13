package handlers

import (
	"github.com/gin-gonic/gin"
     "net/http"
     "starter/pkg/todo"
)

func GetTodo (c *gin.Context) {
     req := c.Request.URL.Query()

     // validation
     if req["id"] == nil {
          c.JSON(http.StatusUnprocessableEntity, gin.H{
               "success": false,
               "message": "Unprocessable entity",
          })       
          return
     }


     var data todo.TodoRepo
     data = &todo.Todo{}

     data.FindId(req["id"][0])

     c.JSON(http.StatusOK, gin.H{
          "success":true,
          "data": data,
     })
}

func StoreTodo (c *gin.Context) {
     req := &todo.Todo{}
     if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}

     var data todo.TodoRepo
     data = req

     data.Insert()

     c.JSON(http.StatusOK, gin.H{
          "success":true,
          "data": data, 
     })
}
