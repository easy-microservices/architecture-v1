package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Category model definition
type Category struct {
	ID             bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name           string        `json:"name" bson:"name"`
	CreatedAt      time.Time     `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at,omitempty" bson:"updated_at"`
	SubCategorries []Category    `json:"sub_categories" bson:"sub_categories"`
	Products       []Product     `json:"products" bson:"products"`
}
