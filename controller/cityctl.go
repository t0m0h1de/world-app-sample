package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/t0m0h1de/world-app-sample/repository"
)

func CityController(r *gin.Engine, db repository.Database) {
	r.GET("/city/:cityname", func(c *gin.Context) {
		city := c.Param("cityname")
		population, err := db.FindCityPopulation(city)
		if err == nil {
			c.JSON(200, gin.H{
				"message": strconv.Itoa(population),
			})
		} else {
			notFoundResponse(c, fmt.Sprintf("city: %s", city))
		}
	})
}

func notFoundResponse(c *gin.Context, msg string) {
	c.JSON(404, gin.H{
		"message": fmt.Sprintf("Resouce is not found (%s)", msg),
	})
}
