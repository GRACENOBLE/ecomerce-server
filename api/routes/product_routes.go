package routes

import (
	"context"
	"fmt"
	"log"

	"github.com/GRACENOBLE/kampe-backend/database"
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	r.GET("/products", getProducts)
	r.POST("/products", createProduct)
	r.PUT("/products/:id", updateProduct)
	r.DELETE("/products/:id", deleteProduct)
	// Add more product-related routes here
}

func getProducts(c *gin.Context) {
	db := database.ConnectDatabase()
	defer db.Close()

	rows, err := db.Query(context.Background(), "SELECT id, name FROM users")
	if err != nil {
		log.Fatalf("Failed to query data: %v\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatalf("Failed to scan row: %v\n", err)
		}
		fmt.Printf("ID: %s, Name: %s\n", id, name)
	}

}

func createProduct(c *gin.Context) {
	// Handler logic for creating a product
}

func updateProduct(c *gin.Context) {
	// Handler logic for updating a product
}

func deleteProduct(c *gin.Context) {
	// Handler logic for deleting a product
}