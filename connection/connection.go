package connection

import (
	"errors"
	"log"
	"time"

	"github.com/arturoverbel/microservice_compra/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Repository ...
type Repository struct{}

// INFO - to connect mongo
var INFO = &mgo.DialInfo{
	Addrs:    []string{"127.0.0.1:27017"},
	Timeout:  60 * time.Second,
	Database: "cool_db",
	Username: "admin",
	Password: "secret_password",
}

// DBNAME the name of the DB instance
const DBNAME = "cool_db"

// DOCNAME the name of the document
const DOCNAME = "shoppings"

var db *mgo.Database

// COLLECTION - name collection on Mongo
const (
	COLLECTION = "shoppings"
)

// Insert - Insert a Shopping
func Insert(shopping model.Shopping) error {
	session, err := mgo.DialWithInfo(INFO)
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

	session, err := mgo.DialWithInfo(INFO)
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
	session, err := mgo.DialWithInfo(INFO)
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
	session, err := mgo.DialWithInfo(INFO)
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
	session, err := mgo.DialWithInfo(INFO)
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
