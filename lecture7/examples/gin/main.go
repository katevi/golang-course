package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router (with default middleware)
	r := gin.Default()

	// Define routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Gin server!",
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Guest")
		c.String(http.StatusOK, "Hello %s!", name)
	})

	r.POST("/greet", func(c *gin.Context) {
		var json struct {
			Name string `json:"name" binding:"required"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"greeting": "Hello " + json.Name + "!",
		})
	})

	// Start the server on port 8080
	r.Run(":8080") // listen and serve
}
