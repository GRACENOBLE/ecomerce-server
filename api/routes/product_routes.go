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

	rows, err := db.Query(context.Background(), "SELECT p.id AS product_id, p.name AS product_name, p.description, p.base_price, p.sku AS product_sku, p.inventory_quantity AS product_inventory, p.has_variants, v.id AS variant_id, v.sku AS variant_sku, v.price AS variant_price, v.color AS variant_color, v.size AS variant_size, v.inventory_quantity AS variant_inventory, img.id AS image_id, img.url AS image_url, img.thumbnail_url, img.type AS image_type FROM products p LEFT JOIN product_variants v ON v.product_id = p.id LEFT JOIN product_images img ON img.product_id = p.id OR img.variant_id = v.id ORDER BY p.id, v.id, img.position")
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