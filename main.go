package main

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/t0m0h1de/world-app-sample/controller"
	"github.com/t0m0h1de/world-app-sample/errhandler"
	"github.com/t0m0h1de/world-app-sample/repository"
)

func main() {
	c := configure()
	r := gin.Default()
	db := &repository.PostgreSQL{}
	err := db.Connect(c.dbuser, c.dbpassword, c.dbhost, c.dbname, c.dbport, c.dbsslmode)
	errhandler.ErrHandler(err)
	err = controller.CityController(r, db)
	errhandler.ErrHandler(err)
	r.Run((":7879"))
}

type config struct {
	dbuser     string
	dbpassword string
	dbhost     string
	dbname     string
	dbport     int
	dbsslmode  bool
}

func configure() config {
	dbuser := os.Getenv("DBUSER")
	if dbuser == "" {
		dbuser = "postgres"
	}
	dbpassword := os.Getenv("DBPASSWORD")
	if dbpassword == "" {
		dbpassword = "password"
	}
	dbhost := os.Getenv("DBHOST")
	if dbhost == "" {
		dbhost = "localhost"
	}
	dbname := os.Getenv("DBNAME")
	if dbname == "" {
		dbname = "postgres"
	}
	dbportstr := os.Getenv("DBPORT")
	dbport, err := strconv.Atoi(dbportstr)
	if err != nil {
		dbport = 30451
	}
	dbsslmodestr := os.Getenv("DBSSLMODE")
	var dbsslmode bool = true
	if dbsslmodestr == "" || dbsslmodestr == "false" || dbsslmodestr == "FALSE" {
		dbsslmode = false
	}
	c := config{
		dbuser:     dbuser,
		dbpassword: dbpassword,
		dbhost:     dbhost,
		dbname:     dbname,
		dbport:     dbport,
		dbsslmode:  dbsslmode,
	}
	return c
}
