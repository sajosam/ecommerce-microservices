package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	users := map[string]gin.H{
		"1": {"id": 1, "name": "David-User"},
		"2": {"id": 2, "name": "Ajce-User"},
	}

	r.GET("/users/:id", func(c *gin.Context) {
		userID := c.Param("id")
		if user, exists := users[userID]; exists {
			c.JSON(200, user)
		} else {
			c.JSON(404, gin.H{"error": "User not found"})
		}
	})

	r.Run(":5001")
}
