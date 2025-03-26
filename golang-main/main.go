package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func fetchData(url string) interface{} {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return gin.H{"error": "Failed to fetch data"}
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var data interface{}
	json.Unmarshal(body, &data)
	return data
}

func main() {
	r := gin.Default()

	r.GET("/users/:id", func(c *gin.Context) {
		userID := c.Param("id")
		c.JSON(200, fetchData("http://golang-user:5001/users/"+userID))
	})

	r.GET("/orders/:id", func(c *gin.Context) {
		orderID := c.Param("id")
		c.JSON(200, fetchData("http://fastapi-order:5002/orders/"+orderID))
	})

	r.GET("/payments/:id", func(c *gin.Context) {
		paymentID := c.Param("id")
		c.JSON(200, fetchData("http://flask-payment:5003/payments/"+paymentID))
	})

	r.GET("/products/:id", func(c *gin.Context) {
		productID := c.Param("id")
		c.JSON(200, fetchData("http://nodejs-product:5004/products/"+productID))
	})

	r.GET("/all-details/:user_id", func(c *gin.Context) {
		userID := c.Param("user_id")
		c.JSON(200, gin.H{
			"user":     fetchData("http://golang-user:5001/users/" + userID),
			"orders":   fetchData("http://fastapi-order:5002/orders/" + userID),
			"payments": fetchData("http://flask-payment:5003/payments/" + userID),
			"products": fetchData("http://nodejs-product:5004/products/" + userID),
		})
	})

	r.Run(":5000")
}
