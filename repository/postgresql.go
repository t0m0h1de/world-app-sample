package repository

import (
	"database/sql"
	"fmt"
)

type PostgreSQL struct {
	dsn string
	db  *sql.DB
}

func (pg *PostgreSQL) Connect(user string, password string, host string, dbname string, port int, sslmode bool) error {
	var sslcap string
	if sslmode {
		sslcap = "enable"
	} else {
		sslcap = "disable"
	}
	pg.dsn = fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d sslmode=%s", user, password, host, dbname, port, sslcap)
	var err error
	pg.db, err = sql.Open("postgres", pg.dsn)
	return err
}

func (pg PostgreSQL) Close() {
	pg.db.Close()
}

func (pg *PostgreSQL) FindPopulation(city string) (int, error) {
	query := fmt.Sprintf("SELECT population FROM city WHERE name = '%s'", city)
	var population int
	err := pg.db.QueryRow(query).Scan(&population)
	return population, err
}
