package product

import (
	"time"

	"github.com/easyms/archv1/pc-service/internal/db"
	"github.com/easyms/archv1/pc-service/internal/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MgoProductRepository defines the persistance layer for
// manipulating product information through mongodb.
type MgoProductRepository struct {
	conn       *db.Connection
	collection *mgo.Collection
}

// NewMgoProductRepository creates a new mongo product repository
func NewMgoProductRepository(c *db.Connection) *MgoProductRepository {

	return &MgoProductRepository{
		conn:       c,
		collection: c.Use("products"),
	}
}

// Create method creates an new product in the collection
func (repo *MgoProductRepository) Create(p *models.Product) (int64, error) {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	err := repo.collection.Insert(p)
	if err != nil {
		return 0, err
	}
	return 0, err
}

// GetByID fetches a product by it's ID
func (repo *MgoProductRepository) GetByID(id string) (*models.Product, error) {

	product := models.Product{}

	err := repo.collection.FindId(bson.ObjectIdHex(id)).One(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// GetByName fetches a product by it's Name
func (repo *MgoProductRepository) GetByName(name string) (*models.Product, error) {

	product := models.Product{}

	err := repo.collection.Find(bson.M{"name": name}).One(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// All retrieves all products stored in the collection
func (repo *MgoProductRepository) All() ([]*models.Product, error) {
	products := []*models.Product{}
	err := repo.collection.Find(nil).All(&products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// Update method, updates a specific product in the collection
func (repo *MgoProductRepository) Update(p *models.Product) (*models.Product, error) {
	p.UpdatedAt = time.Now()
	err := repo.collection.UpdateId(p.ID, &p)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Delete method, removes a specific product from the collection
func (repo *MgoProductRepository) Delete(id string) (bool, error) {
	err := repo.collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		return false, err
	}
	return true, nil
}
