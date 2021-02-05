package controllers

import (
	"fmt"
	"net/http"

	"github.com/easyms/archv1/pc-service/internal/models"
	repo "github.com/easyms/archv1/pc-service/internal/repository/product"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// ProductController enumerates the product controller methods
type ProductController struct {
	repository *repo.MgoProductRepository
	logger     *log.Logger
}

// NewProductController creates a new product controller
func NewProductController(r *repo.MgoProductRepository, l *log.Logger) *ProductController {

	return &ProductController{
		repository: r,
		logger:     l,
	}
}

// CreateProduct creates a new product
func (p *ProductController) CreateProduct(c *gin.Context) {

	var product models.Product
	c.BindJSON(&product)

	id, err := p.repository.Create(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error has occured: %s", err))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"_id": id})
}

// GetAllProducts retrieves all products
func (p *ProductController) GetAllProducts(c *gin.Context) {

	products, err := p.repository.All()
	if err != nil {
		c.JSON(http.StatusNotFound,
			fmt.Sprintf("Error occured while retrieving products: %s", err))
		return
	}
	c.JSON(http.StatusOK, products)

}

// GetProductByID retrieves product by its ID
func (p *ProductController) GetProductByID(c *gin.Context) {

	id := c.Param("id")

	product, err := p.repository.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("No prodct found for _id: %s", id))
	}
	c.JSON(http.StatusOK, product)
}

// GetProductByName retrieves product by its name
func (p *ProductController) GetProductByName(c *gin.Context) {

	name := c.Param("name")

	product, err := p.repository.GetByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("No product found for name: %s", name))
		return
	}
	c.JSON(http.StatusOK, product)
}

// UpdateProduct updates product
func (p *ProductController) UpdateProduct(c *gin.Context) {
	var productUpdater models.Product
	c.BindJSON(&productUpdater)

	_, err := p.repository.Update(&productUpdater)
	if err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("Error has occured: %s", err))
		return
	}
	c.JSON(http.StatusOK, "Product succesfully updated")
}

// DeleteProduct deletes product
func (p *ProductController) DeleteProduct(c *gin.Context) {

	id := c.Param("id")

	_, err := p.repository.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("Error has occured: %s", err))
		return
	}
	c.JSON(http.StatusOK, "Product succesfully deleted")
}
