package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cityCtrl()
}

func errHandler(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func cityCtrl() {
	// Ginエンジンのインスタンス作成
	r := gin.Default()

	r.GET("/city/:cityname", func(c *gin.Context) {
		cityname := c.Param("cityname")
		populationMsg := fmt.Sprintf("populatoin: %d", citySvc(cityname))
		c.JSON(200, gin.H{
			"message": populationMsg,
		})
	})

	r.Run((":7879"))
}

func citySvc(city string) int {
	// db接続
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=postgres port=5433 sslmode=disable")
	errHandler(err)

	population := findPopulation(db, city)
	return population
}

func findPopulation(db *sql.DB, city string) int {
	// データの検索
	query := fmt.Sprintf("SELECT population FROM city WHERE name = '%s'", city)
	var population int
	err := db.QueryRow(query).Scan(&population)
	errHandler(err)
	return population
}
