package repository

import "github.com/t0m0h1de/world-app-sample/model"

type Database interface {
	Connect(string, string, string, string, int, bool)
	Close()
	FindCity(cityname string) (model.City, error)
	FindCityById(id int) (model.City, error)
}
