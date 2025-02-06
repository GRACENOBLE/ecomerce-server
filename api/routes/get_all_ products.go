package routes

import (
	"context"
	"log"

	"github.com/GRACENOBLE/kampe-backend/database"
	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	db := database.ConnectDatabase()
	defer db.Close()

	rows, err := db.Query(context.Background(), "SELECT DISTINCT ON (V.ID) P.ID AS PRODUCT_ID, P.NAME AS PRODUCT_NAME, P.DESCRIPTION, P.BASE_PRICE, P.SKU AS PRODUCT_SKU, P.INVENTORY_QUANTITY AS PRODUCT_INVENTORY, P.HAS_VARIANTS, P.TAGS, P.RATING AS PRODUCT_RATING, P.VERIFIED_PURCHASES, P.DISCOUNT, V.ID AS VARIANT_ID, V.SKU AS VARIANT_SKU, V.PRICE AS VARIANT_PRICE, V.COLOR AS VARIANT_COLOR, V.SIZE AS VARIANT_SIZE, V.INVENTORY_QUANTITY AS VARIANT_INVENTORY, IMG.ID AS IMAGE_ID, IMG.URL AS IMAGE_URL, IMG.THUMBNAIL_URL, IMG.TYPE AS IMAGE_TYPE FROM PRODUCTS P LEFT JOIN PRODUCT_VARIANTS V ON V.PRODUCT_ID = P.ID LEFT JOIN PRODUCT_IMAGES IMG ON IMG.PRODUCT_ID = P.ID OR (IMG.VARIANT_ID = V.ID AND P.HAS_VARIANTS = TRUE) ORDER BY V.ID, IMG.POSITION;")
	if err != nil {
		log.Fatalf("Failed to query data: %v\n", err)
	}
	defer rows.Close()

	productsMap := make(map[string]map[string]interface{})

	for rows.Next() {
		var (
			productID         string
			productName       string
			description       string
			basePrice         float64
			productSKU        string
			productInventory  int
			hasVariants       bool
			tags              []string
			productRating     float64
			verifiedPurchases int64
			discount          int64
			variantID         *string
			variantSKU        *string
			variantPrice      *float64
			variantColor      *string
			variantSize       *string
			variantInventory  *int
			imageID           *string
			imageURL          *string
			thumbnailURL      *string
			imageType         *string
		)
		err := rows.Scan(
			&productID, &productName, &description, &basePrice, &productSKU, &productInventory, &hasVariants, &tags, &productRating, &verifiedPurchases, &discount,
			&variantID, &variantSKU, &variantPrice, &variantColor, &variantSize, &variantInventory,
			&imageID, &imageURL, &thumbnailURL, &imageType,
		)
		if err != nil {
			log.Fatalf("Failed to scan row: %v\n", err)
		}

		if _, exists := productsMap[productID]; !exists {
			productsMap[productID] = map[string]interface{}{
				"id":                 productID,
				"name":               productName,
				"description":        description,
				"base_price":         basePrice,
				"sku":                productSKU,
				"inventory":          productInventory,
				"has_variants":       hasVariants,
				"tags":               tags,
				"product_rating":     productRating,
				"verified_purchases": verifiedPurchases,
				"discount": discount,
				"variants":           []map[string]interface{}{},
				"images":             []map[string]interface{}{},
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
			productsMap[productID]["variants"] = append(productsMap[productID]["variants"].([]map[string]interface{}), variant)

			if imageID != nil {
				image := map[string]interface{}{
					"id":            *imageID,
					"url":           *imageURL,
					"thumbnail_url": *thumbnailURL,
					"type":          *imageType,
				}
				variant["images"] = append(variant["images"].([]map[string]interface{}), image)
			}
		} else if imageID != nil {
			image := map[string]interface{}{
				"id":            *imageID,
				"url":           *imageURL,
				"thumbnail_url": *thumbnailURL,
				"type":          *imageType,
			}
			productsMap[productID]["images"] = append(productsMap[productID]["images"].([]map[string]interface{}), image)
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