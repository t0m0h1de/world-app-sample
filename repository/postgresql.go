package repository

import (
	"database/sql"
	"fmt"

	"github.com/t0m0h1de/world-app-sample/errhandler"
)

type PostgreSQL struct {
	dsn string
	db  *sql.DB
}

func (pg *PostgreSQL) Connect(user string, password string, host string, dbname string, port int, sslmode bool) {
	var sslcap string
	if sslmode {
		sslcap = "enable"
	} else {
		sslcap = "disable"
	}
	pg.dsn = fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d sslmode=%s", user, password, host, dbname, port, sslcap)
	var err error
	pg.db, err = sql.Open("postgres", pg.dsn)
	errhandler.ErrHandler(err)
	err = pg.db.Ping()
	errhandler.ErrHandler(err)
}

func (pg PostgreSQL) Close() {
	pg.db.Close()
}

func (pg *PostgreSQL) FindCityPopulation(city string) (int, error) {
	err := pg.db.Ping()
	errhandler.ErrHandler(err)
	query := fmt.Sprintf("SELECT population FROM city WHERE name = '%s'", city)
	var population int
	err = pg.db.QueryRow(query).Scan(&population)
	return population, err
}
