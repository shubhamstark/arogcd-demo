package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Order struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
}

var orders = []Order{
	{ID: "1", Amount: 100},
	{ID: "2", Amount: 200},
}

func main() {
	r := gin.Default()

	r.GET("/orders", getOrders)
	r.POST("/orders", createOrder)

	r.Run(":8080") // Run on port 8080
}

func getOrders(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, orders)
}

func createOrder(c *gin.Context) {
	var newOrder Order

	if err := c.BindJSON(&newOrder); err != nil {
		return
	}

	orders = append(orders, newOrder)
	c.IndentedJSON(http.StatusCreated, newOrder)
}
