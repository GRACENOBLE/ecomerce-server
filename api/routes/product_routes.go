package routes

import (
	"context"
	// "fmt"
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

	productsMap := make(map[string]map[string]interface{})

	for rows.Next() {
		var (
			productID        string
			productName      string
			description      string
			basePrice        float64
			productSKU       string
			productInventory int
			hasVariants      bool
			variantID        *string
			variantSKU       *string
			variantPrice     *float64
			variantColor     *string
			variantSize      *string
			variantInventory *int
			imageID          *string
			imageURL         *string
			thumbnailURL     *string
			imageType        *string
		)
		err := rows.Scan(
			&productID, &productName, &description, &basePrice, &productSKU, &productInventory, &hasVariants,
			&variantID, &variantSKU, &variantPrice, &variantColor, &variantSize, &variantInventory,
			&imageID, &imageURL, &thumbnailURL, &imageType,
		)
		if err != nil {
			log.Fatalf("Failed to scan row: %v\n", err)
		}

		if _, exists := productsMap[productID]; !exists {
			productsMap[productID] = map[string]interface{}{
				"id":          productID,
				"name":        productName,
				"description": description,
				"base_price":  basePrice,
				"sku":         productSKU,
				"inventory":   productInventory,
				"has_variants": hasVariants,
				"variants":    []map[string]interface{}{},
			}
		}

		if variantID != nil {
			variant := map[string]interface{}{
				"id":        *variantID,
				"sku":       *variantSKU,
				"price":     *variantPrice,
				"color":     *variantColor,
				"size":      *variantSize,
				"inventory": *variantInventory,
				"images":    []map[string]interface{}{},
			}

			if imageID != nil {
				image := map[string]interface{}{
					"id":            *imageID,
					"url":           *imageURL,
					"thumbnail_url": *thumbnailURL,
					"type":          *imageType,
				}
				variant["images"] = append(variant["images"].([]map[string]interface{}), image)
			}

			productsMap[productID]["variants"] = append(productsMap[productID]["variants"].([]map[string]interface{}), variant)
		}
	}

	products := []map[string]interface{}{}
	for _, product := range productsMap {
		products = append(products, product)
	}

	c.JSON(200, gin.H{
		"products": products,
	})
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
