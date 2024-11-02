package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/t0m0h1de/world-app-sample/repository"
)

func CityController(r *gin.Engine, db repository.Database) {
	r.GET("/city/:cityname", func(c *gin.Context) {
		cityname := c.Param("cityname")
		city, err := db.FindCity(cityname)
		if err == nil {
			c.JSON(200, gin.H{
				"id":          city.Id,
				"name":        city.Name,
				"countrycode": city.CountryCode,
				"district":    city.District,
				"population":  city.Population,
			})
		} else {
			notFoundResponse(c, fmt.Sprintf("cityname: %s", cityname))
		}
	})

	r.GET("/city/id/:id", func(c *gin.Context) {
		inputId := c.Param("id")
		id, err := strconv.Atoi(inputId)
		if err != nil {
			badRequestResponse(c, fmt.Sprintf("id: %s must be numeric", inputId))
		}
		city, err := db.FindCityById(id)
		if err == nil {
			c.JSON(200, gin.H{
				"id":          city.Id,
				"name":        city.Name,
				"countrycode": city.CountryCode,
				"district":    city.District,
				"population":  city.Population,
			})
		} else {
			notFoundResponse(c, fmt.Sprintf("id: %s", inputId))
		}
	})
}

func notFoundResponse(c *gin.Context, msg string) {
	c.JSON(404, gin.H{
		"message": fmt.Sprintf("Resouce is not found (%s)", msg),
	})
}

func badRequestResponse(c *gin.Context, msg string) {
	c.JSON(400, gin.H{
		"message": fmt.Sprintf("Bad request (%s)", msg),
	})
}
