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

func GetPlanet(c *gin.Context) error {
	c.JSON(200, gin.H{
		"message": "Teste Hello",
	})
	return nil
}

func UpdatePlanet(c *gin.Context) error {
	return nil
}

func DeletePlanet(c *gin.Context) error {
	return nil
}