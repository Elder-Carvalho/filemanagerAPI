package routes

import (
	"database/sql"
	"filemanagerAPI/controllers"
	"github.com/labstack/echo"
)

type Router struct{}

func (r Router) SetupRoutes(e *echo.Echo, db *sql.DB) {

	// ac := controllers.AuthController{DB: db}
	// e.POST("/login", ac.Login)

	uc := controllers.UserController{DB: db}
	e.GET("/users", uc.FindAll)
	e.POST("/users", uc.Insert)
}
