package handlers

import (
	"net/http"
	"starter/platform/newsfeed"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func GetNewsFeed(c *gin.Context) {
	var data newsfeed.NewsFeedRepo

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"success": false,
			"message": "Unprocessable entity",
		})
		return
	}

	if err := validator.New().Struct(&data); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": err.Field() + " is " + err.ActualTag() + " " + err.Type().Name(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
