// Package classification of Category API
//
// Documentation for Category API
//
// Schemes: http
// BasePath: /categories
// Version: 1.0.0
//
// Consumes:
//	- application/json
//
// Produces:
//	- application/json
// swagger:meta
package controllers

import (
	"fmt"
	"net/http"

	"github.com/easyms/archv1/pc-service/internal/models"
	repo "github.com/easyms/archv1/pc-service/internal/repository/category"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// CategoryController enumerates the category controller methods
type CategoryController struct {
	repository *repo.MgoCategoryRepository
	logger     *log.Logger
}

// NewCategoryController create a new category controller
func NewCategoryController(r *repo.MgoCategoryRepository, l *log.Logger) *CategoryController {

	return &CategoryController{
		repository: r,
		logger:     l,
	}
}

// CreateCategory creates an new category
func (cat *CategoryController) CreateCategory(c *gin.Context) {

	var category models.Category
	c.BindJSON(&category)

	id, err := cat.repository.Create(&category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error has occured: %s", err))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"_id": id})
}

// GetAllCategories retrieve all categories
func (cat *CategoryController) GetAllCategories(c *gin.Context) {

	categories, err := cat.repository.All()
	if err != nil {
		c.JSON(http.StatusNotFound,
			fmt.Sprintf("Error occured while retrieving categories: %s", err))
		return
	}
	c.JSON(http.StatusOK, categories)
}

// GetCategoryByName retreive category by name
func (cat *CategoryController) GetCategoryByName(c *gin.Context) {

	name := c.Param("name")

	cat.logger.Info("GetCategoryByName() - searching for category with name:", name)

	category, err := cat.repository.GetByName(name)

	if err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("No category found for name: %s", name))
		return
	}

	c.JSON(http.StatusOK, category)
}

// AddSubcategory creates and adds a new subcategory to category
func (cat *CategoryController) AddSubcategory(c *gin.Context) {

	name := c.Param("name")

	var (
		subcategory models.Category
		category    *models.Category
		err         error
	)

	c.BindJSON(&subcategory)

	category, err = cat.repository.GetByName(name)

	if err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("No category found for name: %s", name))
		return
	}

	category.SubCategorries = append(category.SubCategorries, subcategory)
	_, err = cat.repository.Update(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error has occured: %s", err))
		return
	}

	cat.logger.Info(fmt.Sprintf("AddSubcategory() - Added subcategory: %s under category: %s",
		subcategory.Name, category.Name))
	c.JSON(http.StatusCreated, subcategory)

}

// GetAllSubcategories retrevies all subcategories of a category
func (cat *CategoryController) GetAllSubcategories(c *gin.Context) {
	name := c.Param("name")

	category, err := cat.repository.GetByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("No category found for name: %s", name))
		return
	}
	c.JSON(http.StatusOK, category.SubCategorries)
}

// UpdateCategory update category
func (cat *CategoryController) UpdateCategory(c *gin.Context) {
	var categoryUpdater models.Category
	c.BindJSON(&categoryUpdater)

	_, err := cat.repository.Update(&categoryUpdater)
	if err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("Error has occured: %s", err))
		return
	}
	c.JSON(http.StatusOK, "Category succesfully updated")
}

// DeleteCategory delete category
func (cat *CategoryController) DeleteCategory(c *gin.Context) {

	id := c.Param("id")

	_, err := cat.repository.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("Error has occured: %s", err))
		return
	}
	c.JSON(http.StatusOK, "Category succesfully deleted")
}
