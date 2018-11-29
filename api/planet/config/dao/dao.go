package dao

import (
	"log"

	"api_rest/api/planet/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DAO struct {
	Server   string
	Database string
}

var db *mgo.Database

func (m *DAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// func (m *DAO) GetAll() ([]model.Planet, error) {
// 	var planets []model.Planet
// 	err := db.C("planets").Find(bson.M{}).All(&planets)
// 	return planets, err
// }

// func (m *DAO) GetByID(id string) (model.Planet, error) {
// 	var planet model.Planet
// 	err := db.C("planets").FindId(bson.ObjectIdHex(id)).One(&planet)
// 	return planet, err
// }

// func (m *DAO) Create(planet model.Planet) error {
// 	err := db.C("planets").Insert(&planet)
// 	return err
// }

// func (m *DAO) Delete(id string) error {
// 	err := db.C("planets").RemoveId(bson.ObjectIdHex(id))
// 	return err
// }

// func (m *DAO) Update(id string, planet model.Planet) error {
// 	err := db.C("planets").UpdateId(bson.ObjectIdHex(id), &planet)
// 	return err
// }