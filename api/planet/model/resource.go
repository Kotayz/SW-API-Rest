package model

import (
	"encoding/json"
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
	pipeString := `[{"$match":{"username":` + planetName + `}}]` //Utilizado '$match' para evitar problema de uppercase ou lower case
	pipe := []bson.M{}
	err := json.Unmarshal([]byte(pipeString), &pipe)
	err = db.C("planets").Pipe(pipe).One(&planet)

	return planet, err
}

func (m *DAO) Create(planet *Planet) error {
	planet.ID = bson.NewObjectId()
	err := db.C("planets").Insert(planet)
	return err
}

func (m *DAO) Delete(id string) error {
	err := db.C("planets").RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *DAO) Update(id string, planet Planet) error {
	err := db.C("planets").UpdateId(bson.ObjectIdHex(id), &planet)
	return err
}
