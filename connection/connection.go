package connection

import (
	"log"
	"errors"

	"github.com/arturoverbel/microservice_compra/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Repository ...
type Repository struct{}

// SERVER the DB server
const SERVER = "localhost:27017"

// DBNAME the name of the DB instance
const DBNAME = "pos"

// DOCNAME the name of the document
const DOCNAME = "shoppings"
var db *mgo.Database

// COLLECTION - name collection on Mongo
const (
	COLLECTION = "shoppings"
)

// Insert - Insert a Shopping
func Insert(shopping model.Shopping) error {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	shopping.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(shopping)

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// FindByID - ...
func FindByID(id string) (model.Shopping, error) {
	var shopping model.Shopping
	if !bson.IsObjectIdHex(id) {
		err := errors.New("Invalid ID")
		return shopping, err
	}

	session, err := mgo.Dial(SERVER)
	if err != nil {
		log.Fatal(err)
		return shopping, err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)

	oid := bson.ObjectIdHex(id)
	err = c.FindId(oid).One(&shopping)
	return shopping, err
}

// Update - ..
func Update(shopping model.Shopping) error {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		log.Fatal(err)
		return err
	}
 	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	err = c.UpdateId(shopping.ID, &shopping)
	return err
}

// FindByUser - ...
func FindByUser(idUser int) ([]model.Shopping, error) {
	var shoppings []model.Shopping
	session, err := mgo.Dial(SERVER)
	if err != nil {
		log.Fatal(err)
		return shoppings, err
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)

	err = c.Find(bson.M{"user": idUser}).All(&shoppings)
	return shoppings, err
}

// Delete - ...
func Delete(id string) error {
	if !bson.IsObjectIdHex(id) {
		err := errors.New("Invalid ID")
		return err
	}
	session, err := mgo.Dial(SERVER)
	if err != nil {
		log.Fatal(err)
		return err
	}
 	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)

	oid := bson.ObjectIdHex(id)
	err = c.RemoveId(oid)
	return err
}
