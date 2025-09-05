package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	users := map[string]gin.H{
		"1": {"id": 1, "name": "David"},
		"2": {"id": 2, "name": "Ajce"},
	}

	r.GET("/users/:id", func(c *gin.Context) {
		userID := c.Param("id")
		if user, exists := users[userID]; exists {
			c.JSON(200, user)
		} else {
			c.JSON(404, gin.H{"error": "User not found"})
		}
	})

	// health check with custom msg
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "User service is running on port 5001"})
	})

	r.Run(":5001")
}
