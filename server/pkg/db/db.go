package db

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

type AsosDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	ItemsCollection      = "items"
	CategoriesCollection = "categories"
)

func (m *AsosDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}
