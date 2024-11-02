package repository

import (
	"database/sql"
	"fmt"

	"github.com/t0m0h1de/world-app-sample/errhandler"
	"github.com/t0m0h1de/world-app-sample/model"
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

func (pg *PostgreSQL) FindCity(cityname string) (model.City, error) {
	err := pg.db.Ping()
	errhandler.ErrHandler(err)
	query := fmt.Sprintf("SELECT * FROM city WHERE name = '%s'", cityname)
	var id int
	var name string
	var countrycode string
	var district string
	var population int
	err = pg.db.QueryRow(query).Scan(&id, &name, &countrycode, &district, &population)
	city := model.City{
		Id:          id,
		Name:        name,
		CountryCode: countrycode,
		District:    district,
		Population:  population,
	}
	return city, err
}

func (pg *PostgreSQL) FindCityById(id int) (model.City, error) {
	err := pg.db.Ping()
	errhandler.ErrHandler(err)
	query := fmt.Sprintf("SELECT * FROM city WHERE id = %d", id)
	var name string
	var countrycode string
	var district string
	var population int
	err = pg.db.QueryRow(query).Scan(&id, &name, &countrycode, &district, &population)
	city := model.City{
		Id:          id,
		Name:        name,
		CountryCode: countrycode,
		District:    district,
		Population:  population,
	}
	return city, err
}
