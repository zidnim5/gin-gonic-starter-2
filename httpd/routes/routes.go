package routes

import "github.com/gin-gonic/gin"

func Routes(route *gin.Engine) {

	/** TODO separete routes */
	todoRoute(route)
	newsfeedRoute(route)

}
