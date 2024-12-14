package main

import (
	"github.com/gin-gonic/gin"
)

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var Products []Product

func main() {
	r := gin.Default()

	r.GET("/products", func(c *gin.Context) {
		c.JSON(200, Products)
	})

	r.POST("/products", func(c *gin.Context) {
		var newProduct Product
		if err := c.ShouldBindJSON(&newProduct); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		Products = append(Products, newProduct)
		c.JSON(201, gin.H{"message": "Product created successfully"})
	})

	r.PUT("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedProduct Product
		if err := c.ShouldBindJSON(&updatedProduct); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		for i, product := range Products {
			if product.ID == id {
				Products[i] = updatedProduct
				c.JSON(200, gin.H{"message": "Product updated successfully"})
				return
			}
		}

		c.JSON(404, gin.H{"error": "Product not found"})
	})

	r.Run(":3000")
}
