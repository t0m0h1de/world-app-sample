package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=postgres port=5433 sslmode=disable")
	errHandler(err)

	// データの検索
	rows, err := db.Query("SELECT name, id, population FROM city")
	errHandler(err)

	for rows.Next() {
		var name string
		var id int
		var population int
		err = rows.Scan(&name, &id, &population)
		errHandler(err)
		fmt.Println(name)
		fmt.Println(id)
		fmt.Println(population)
	}
}

func errHandler(err error) {
	if err != nil {
		panic(err)
	}
}
