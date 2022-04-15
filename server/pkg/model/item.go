package model

import "gopkg.in/mgo.v2/bson"

type Item struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Name     string        `bson:"name" json:"name"`
	Color    string        `bson:"color" json:"color"`
	Size     int           `bson:"size" json:"int"`
	Price    float32       `bson:"price" json:"price"`
	Gender   bool          `bson:"gender" json:"gender"`
	InStock  int           `bson:"in_stock" json:"in_stock"`
	Category string        `bson:"category" json:"category"`
	Brand    string        `bson:"brand" json:"brand"`
}
