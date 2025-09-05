package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func fetchData(url string) interface{} {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching API data:", err)
		return gin.H{"error": "Failed to fetch API Data"}
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var data interface{}
	json.Unmarshal(body, &data)
	return data
}

func main() {
	orderService := os.Getenv("ORDER_SERVICE")
	r := gin.Default()

	// Instead of calling all services, main now delegates to Order service
	r.GET("/users/:id", func(c *gin.Context) {
		userID := c.Param("id")
		c.JSON(200, fetchData("http://"+orderService+"/orders/"+userID))
	})

	// For debugging - you can still call each service directly if needed
	r.GET("/orders/:id", func(c *gin.Context) {
		orderID := c.Param("id")
		c.JSON(200, fetchData("http://"+orderService+"/orders/"+orderID))
	})

	r.GET("/payments/:id", func(c *gin.Context) {
		paymentID := c.Param("id")
		paymentService := os.Getenv("PAYMENT_SERVICE")
		c.JSON(200, fetchData("http://"+paymentService+"/payments/"+paymentID))
	})

	r.GET("/products/:id", func(c *gin.Context) {
		productID := c.Param("id")
		productService := os.Getenv("PRODUCT_SERVICE")
		c.JSON(200, fetchData("http://"+productService+"/products/"+productID))
	})

	r.Run(":5000")
}
