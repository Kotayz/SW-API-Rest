package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"gopkg.in/mgo.v2/bson"
)

type Planet struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Nome      string        `bson:"nome" json:"nome"`
	Clima     string        `bson:"clima" json:"clima"`
	Terreno   string        `bson:"terreno" json:"terreno"`
	Aparicoes int           `bson:"aparicoes" json:"aparicoes"`
}

type SWAPIResult struct {
	Results []SWAPIPlanet `json:"results"`
}

type SWAPIPlanet struct {
	Name           string   `json:"name"`
	RotationPeriod string   `json:"rotation_period"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Diameter       string   `json:"diameter"`
	Climate        string   `json:"climate"`
	Gravity        string   `json:"gravity"`
	Terrain        string   `json:"terrain"`
	SurfaceWater   string   `json:"surface_water"`
	Population     string   `json:"population"`
	Residents      []string `json:"residents"`
	Films          []string `json:"films"`
	Created        string   `json:"created"`
	Edited         string   `json:"edited"`
	URL            string   `json:"url"`
}

func (p *Planet) Save() error {
	err := p.Validate()
	if err != nil {
		return err
	}

	aparicoes, err := GetPlanetRequest(p.Nome)
	if err != nil {
		return err
	}

	p.Aparicoes = aparicoes
	createErr := resource.Create(p)
	if createErr != nil {
		return createErr
	}

	return nil
}

func (Planet) Get(id string) (Planet, error) {
	planet, err := resource.GetByID(id)
	if err != nil {
		return Planet{}, err
	}

	return planet, nil
}

func (Planet) GetByName(planetName string) (Planet, error) {
	planet, err := resource.GetByName(planetName)
	if err != nil {
		return Planet{}, err
	}

	return planet, nil
}

func (Planet) GetAll() ([]Planet, error) {
	planets, err := resource.GetAll()
	if err != nil {
		return nil, err
	}

	return planets, nil
}

func (Planet) Delete(id string) error {
	err := resource.Delete(id)
	if err != nil {
		return err
	}

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

func GetPlanetRequest(planetName string) (int, error) {
	name := url.QueryEscape(planetName)
	url := fmt.Sprintf("http://swapi.co/api/planets/?search=%s", name)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
		return 0, readErr
	}

	var result SWAPIResult
	unmarshalErr := json.Unmarshal(body, &result)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
		return 0, unmarshalErr
	}

	if len(result.Results) > 0 {
		if len(result.Results[0].Films) > 0 {
			return len(result.Results[0].Films), nil
		}
		return 0, nil
	}

	return 0, errors.New("Planet not found")
}
