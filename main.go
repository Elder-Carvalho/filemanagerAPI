package main

import (
	"database/sql"
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/kataras/go-sessions"
	"filemanagerAPI/routes"
	"github.com/labstack/echo"
	"log"
)

var db *sql.DB
var err error

func connectDB() {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1)/file_manager")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	connectDB()
	defer db.Close()

	e := echo.New()
	r := routes.Router{}
	r.SetupRoutes(e, db)
	e.Logger.Fatal(e.Start(":3000"))
}
