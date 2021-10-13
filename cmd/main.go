package main

import (
	"starter/cmd/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.Routes(router)

	router.Run(":8090")
}
