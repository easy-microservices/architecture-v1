package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Product model definition
type Product struct {
	ID          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	Type        string        `json:"type" bson:"type"`
	Description string        `json:"description" bson:"description"`
	SKU         string        `json:"_sku" bson:"_sku"`
	Slug        string        `json:"slug" bson:"slug"`

	Price    float64 `json:"price" bson:"price"`
	Quantity uint32  `json:"quantity" bson:"quantity"`

	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}
