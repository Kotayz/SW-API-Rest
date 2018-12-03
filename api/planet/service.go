package planet

import (
	"sw-api-rest/api/planet/model"
)

func CreatePlanet(planet *model.Planet) error {
	return planet.Save()
}

func GetPlanets() ([]model.Planet, error) {
	return model.Planet{}.GetAll()
}

func GetPlanet(id string) (model.Planet, error) {
	return model.Planet{}.Get(id)
}

func GetPlanetByName(planetName string) (model.Planet, error) {
	return model.Planet{}.GetByName(planetName)
}

func DeletePlanet(id string) error {
	return model.Planet{}.Delete(id)
}
