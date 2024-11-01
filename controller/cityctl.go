package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/t0m0h1de/world-app-sample/repository"
)

func CityController(r *gin.Engine, db repository.Database) error {
	var err error
	r.GET("/city/:cityname", func(c *gin.Context) {
		city := c.Param("cityname")
		population, err := db.FindPopulation(city)
		if err == nil {
			c.JSON(200, gin.H{
				"message": strconv.Itoa(population),
			})
		}
	})
	return err
}
