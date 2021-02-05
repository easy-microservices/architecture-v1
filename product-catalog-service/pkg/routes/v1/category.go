package v1

import (
	"github.com/easyms/archv1/pc-service/internal/controllers"
	"github.com/easyms/archv1/pc-service/internal/db"
	repo "github.com/easyms/archv1/pc-service/internal/repository/category"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// ApplyCategoryRoutes adds category routes to the root group
func ApplyCategoryRoutes(r *gin.RouterGroup, dbConn *db.Connection, logger *log.Logger) {

	repository := repo.NewMgoCategoryRepository(dbConn)
	controller := controllers.NewCategoryController(repository, logger)

	category := r.Group("/categories")
	{
		category.GET("/", controller.GetAllCategories)
		category.GET("/:name", controller.GetCategoryByName)
		category.GET("/:name/subcategories", controller.GetAllSubcategories)
		category.POST("/", controller.CreateCategory)
		category.POST("/:name", controller.AddSubcategory)
		category.DELETE("/:id", controller.DeleteCategory)
		category.PUT("/", controller.UpdateCategory)
	}
}
