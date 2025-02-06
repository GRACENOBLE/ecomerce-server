package api

import (
	"github.com/GRACENOBLE/kampe-backend/api/routes"
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	r.GET("/products", routes.GetAllProducts)
	r.GET("/products/filter", routes.GetFilteredProducts)
}
