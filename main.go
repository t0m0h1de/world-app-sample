package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/t0m0h1de/world-app-sample/controller"
	"github.com/t0m0h1de/world-app-sample/errhandler"
	"github.com/t0m0h1de/world-app-sample/repository"
)

func main() {
	r := gin.Default()
	db := &repository.PostgreSQL{}
	err := db.Connect("postgres", "password", "postgres", 30451, false)
	errhandler.ErrHandler(err)
	err = controller.CityController(r, db)
	errhandler.ErrHandler(err)
	r.Run((":7879"))
}
