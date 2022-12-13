package handlers

import "github.com/gin-gonic/gin"

func HomeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"Name":        "Pokemon API",
		"Description": "This is a REST API to get Pokemon information",
		"Version":     "1.0.0",
	})
}
