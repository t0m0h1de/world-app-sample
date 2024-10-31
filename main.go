package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=postgres port=5433 sslmode=disable")
	errHandler(err)
	city := "Ede"
	population := findPopulation(db, city)
	fmt.Printf("The population of %s: %d\n", city, population)

	// データの検索
}

func findPopulation(db *sql.DB, city string) int {
	// データの検索
	query := fmt.Sprintf("SELECT population FROM city WHERE name = '%s'", city)
	var population int
	err := db.QueryRow(query).Scan(&population)
	errHandler(err)
	return population
}

func errHandler(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
