package main

import (
	"starter/httpd/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.Routes(router)

	router.Run(":8080")
}
