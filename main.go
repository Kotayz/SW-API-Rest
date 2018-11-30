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
	router.PUT("/:id", UpdatePlanet)
	router.DELETE("/:id", DeletePlanet)
	router.Run()
}

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

func GetPlanets(c *gin.Context) {
	planets, err := planetAPI.GetPlanets()
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "planets": planets})
}

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

func GetPlanetByName(c *gin.Context) {
	planet, err := planetAPI.GetPlanetByName(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"status": http.StatusInternalServerError, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "planet": planet})
}

func UpdatePlanet(c *gin.Context) {
	err := planetAPI.UpdatePlanet(c)
	if err != nil {

	}
}

func DeletePlanet(c *gin.Context) {
	err := planetAPI.DeletePlanet(c)
	if err != nil {

	}
}
