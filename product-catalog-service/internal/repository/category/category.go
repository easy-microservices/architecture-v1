package category

import "github.com/easyms/archv1/pc-service/internal/models"

// CategoryRepository provides a general set of methods
// for Category manipulatiob to be implemented in a specific
// database repository
type CategoryRepository interface {
	Create(c *models.Category) (int64, error)
	GetById(id string) (*models.Category, error)
	GetByName(name string) (*models.Category, error)
	All() ([]*models.Category, error)
	Update(c *models.Category) (*models.Category, error)
	Delete(id string) (bool, error)
}
