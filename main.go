package main

import (
	"net/http"

	. "sw-api-rest/api/planet/config"
	"sw-api-rest/api/planet/model"

	planetAPI "sw-api-rest/api/planet"

	"github.com/gin-gonic/gin"
)

var dao = model.DAO{}
var config = Config{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	router := gin.Default()
	router.POST("/", CreatePlanet)
	router.GET("/", GetPlanets)
	router.GET("/:id", GetPlanet)
	router.POST("/filter", GetPlanetByName)
	router.DELETE("/:id", DeletePlanet)

	router.Run()
}

// Método para criar planetas
func CreatePlanet(c *gin.Context) {
	var planet model.Planet
	if err := c.BindJSON(&planet); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"status": http.StatusBadRequest, "error": err.Error(), "params": planet})
		return
	}

	if err := planetAPI.CreatePlanet(&planet); err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": err.Error(), "params": planet})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "Planet Created", "planet": planet})
}

// Método para listar todos os planetas
func GetPlanets(c *gin.Context) {
	planets, err := planetAPI.GetPlanets()
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "planets": planets})
}

// Método para retornar um planeta a partir do id
func GetPlanet(c *gin.Context) {
	id := c.Param("id")

	planet, err := planetAPI.GetPlanet(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "planet": planet})
}

// Método para retornar um planeta a partir do nome
func GetPlanetByName(c *gin.Context) {
	var planet model.Planet
	if err := c.BindJSON(&planet); err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"status": http.StatusBadRequest, "error": err.Error(), "params": planet})
		return
	}

	planet, err := planetAPI.GetPlanetByName(planet.Nome)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "planet": planet})
}

// Método para excluir planeta
func DeletePlanet(c *gin.Context) {
	id := c.Param("id")

	if err := planetAPI.DeletePlanet(id); err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Planet deleted"})
}
