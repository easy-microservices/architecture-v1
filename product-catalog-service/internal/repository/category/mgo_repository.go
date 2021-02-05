package category

import (
	"time"

	"github.com/easyms/archv1/pc-service/internal/db"
	"github.com/easyms/archv1/pc-service/internal/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MgoCategoryRepository defines the persistance layer for
// manipulating category information through mongodb.
type MgoCategoryRepository struct {
	conn       *db.Connection
	collection *mgo.Collection
}

// NewMgoCategoryRepository creates a new mongo category repository
func NewMgoCategoryRepository(c *db.Connection) *MgoCategoryRepository {

	return &MgoCategoryRepository{
		conn:       c,
		collection: c.Use("categories"),
	}
}

// Create method creates an new category in the collection
func (repo *MgoCategoryRepository) Create(c *models.Category) (int64, error) {
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	err := repo.collection.Insert(c)
	if err != nil {
		return 0, err
	}
	return 0, err
}

// GetByID fetches a category by it's ID
func (repo *MgoCategoryRepository) GetByID(id string) (*models.Category, error) {

	category := models.Category{}

	err := repo.collection.FindId(bson.ObjectIdHex(id)).One(&category)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// GetByName fetches a category by it's Name
func (repo *MgoCategoryRepository) GetByName(name string) (*models.Category, error) {

	category := models.Category{}

	err := repo.collection.Find(bson.M{"name": name}).One(&category)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// All retrieves all categories stored in the collection
func (repo *MgoCategoryRepository) All() ([]*models.Category, error) {
	categories := []*models.Category{}
	err := repo.collection.Find(nil).All(&categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// Update method, updates a specific category in the collection
func (repo *MgoCategoryRepository) Update(c *models.Category) (*models.Category, error) {

	c.UpdatedAt = time.Now()
	err := repo.collection.UpdateId(c.ID, &c)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Delete method, removes a specific category from the collection
func (repo *MgoCategoryRepository) Delete(id string) (bool, error) {
	err := repo.collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		return false, err
	}
	return true, nil
}
