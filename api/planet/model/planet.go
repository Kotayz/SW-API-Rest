package model

import (
	"errors"

    "gopkg.in/mgo.v2/bson"
)

type Planet struct {
	ID		bson.ObjectId `bson:"_id" json:"id"`
	Nome	string 		  `bson:"nome" json:"nome"`
	Clima	string 		  `bson:"clima" json:"clima"`
	Terreno	string 		  `bson:"terreno" json:"terreno"`
}

func (p *Planet) Save() error {
	err := p.Validate()
	if err != nil {
		return err
	}

	createErr := resource.Create(p)
	if createErr != nil {
		return createErr
	}

	return nil
}

func (Planet) Get() error {
	return nil
}

func (Planet) GetAll() ([]Planet, error) {
	planets, err := resource.GetAll()
	if err != nil {
		return nil, err
	}

	return planets, nil
}

func (Planet) Update() error {
	return nil
}

func (Planet) Delete() error {
	return nil
}

func (p *Planet) Validate() error {
	if p == nil {
		return errors.New("Planet can't be null")
	}

	if p.Nome == "" {
		return errors.New("Planet name can't be null")
	}

	return nil
}