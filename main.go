package main

import (
	"database/sql"
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/kataras/go-sessions"
	"filemanagerAPI/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"os"
)

var db *sql.DB
var err error

func connectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
	db, err = sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@tcp(127.0.0.1)/"+os.Getenv("DB_NAME"))
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
	e.Use(middleware.Logger())
	r := routes.Router{}
	r.SetupRoutes(e, db)
	e.Logger.Fatal(e.Start(":3000"))
}
