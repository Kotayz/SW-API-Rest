package model

import (
	"log"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DAO struct {
	Server   string
	Database string
}

var db *mgo.Database
var resource *DAO

func (m *DAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *DAO) Create(planet *Planet) error {
	planet.ID = bson.NewObjectId()
	err := db.C("planets").Insert(planet)
	return err
}

func (m *DAO) GetAll() ([]Planet, error) {
	var planets []Planet
	err := db.C("planets").Find(bson.M{}).All(&planets)
	return planets, err
}

func (m *DAO) GetByID(id string) (Planet, error) {
	var planet Planet
	err := db.C("planets").FindId(bson.ObjectIdHex(id)).One(&planet)
	return planet, err
}

func (m *DAO) GetByName(planetName string) (Planet, error) {
	var planet Planet
	err := db.C("planets").Find(bson.M{"nome": bson.M{"$regex": bson.RegEx{planetName, "i"}}}).One(&planet)
	return planet, err
}

func (m *DAO) Delete(id string) error {
	err := db.C("planets").RemoveId(bson.ObjectIdHex(id))
	return err
}
