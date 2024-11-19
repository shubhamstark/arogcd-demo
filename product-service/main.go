package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var products = []Product{
	{ID: "1", Name: "Product 1", Price: 50},
	{ID: "2", Name: "Product 2", Price: 150},
}

func main() {
	r := gin.Default()

	// Product routes
	r.GET("/products", getProducts)
	r.POST("/products", createProduct)

	r.Run(":8080") // Run on port 8080
}

func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

func createProduct(c *gin.Context) {
	var newProduct Product

	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	products = append(products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}
