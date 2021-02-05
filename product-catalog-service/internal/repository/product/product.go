package product

import "github.com/easyms/archv1/pc-service/internal/models"

// ProductRepositry provides a general set of methods
// for Product manipulation to be implemented in a specific
// database repository
type ProductRepositry interface {
	Create(p *models.Product) (int64, error)
	GetByID(id string) (*models.Product, error)
	GetByName(name string) (*models.Product, error)
	All() ([]*models.Product, error)
	Update(p *models.Product) (*models.Product, error)
	Delete(id string) (bool, error)
}
