package v1

import (
	"github.com/easyms/archv1/pc-service/internal/controllers"
	"github.com/easyms/archv1/pc-service/internal/db"
	repo "github.com/easyms/archv1/pc-service/internal/repository/product"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// ApplyProductRoutes adds product routes to the root group
func ApplyProductRoutes(r *gin.RouterGroup, dbConn *db.Connection, logger *log.Logger) {

	repository := repo.NewMgoProductRepository(dbConn)
	controller := controllers.NewProductController(repository, logger)

	product := r.Group("/product")
	{
		product.GET("/", controller.GetAllProducts)
		product.GET("/:id", controller.GetProductByID)
		product.POST("/:id", controller.CreateProduct)
		product.DELETE("/:id", controller.DeleteProduct)
		product.PUT("/", controller.UpdateProduct)
	}
}
