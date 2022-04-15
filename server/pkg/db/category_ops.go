package db

import (
	"asos-api/pkg/model"

	"gopkg.in/mgo.v2/bson"
)

func (a *AsosDAO) FindAllCategories() ([]model.Category, error) {
	var categories []model.Category
	err := db.C(CategoriesCollection).Find(bson.M{}).All(&categories)
	return categories, err
}

func (m *AsosDAO) InsertCategory(category model.Category) error {
	err := db.C(CategoriesCollection).Insert(&category)
	return err
}
