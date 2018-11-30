package planet

import (
	"api_rest/api/planet/model"

	"github.com/gin-gonic/gin"
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

func GetPlanetByName(c *gin.Context) (model.Planet, error) {
	return model.Planet{}.GetByName()
}

func UpdatePlanet(c *gin.Context) error {
	return nil
}

func DeletePlanet(c *gin.Context) error {
	return nil
}