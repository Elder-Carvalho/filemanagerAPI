package controllers

import (
	"database/sql"
	"filemanagerAPI/models"
	"filemanagerAPI/repository"
	"filemanagerAPI/structs"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

const (
	defaultErrMsg = "Ocorreu um erro ao processar a requisição."
)

type UserController struct {
	DB *sql.DB
}

func (u UserController) FindAll(c echo.Context) error {
	ur := repository.UserRepository{DB: u.DB}
	users, err := ur.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, structs.ErrorReponse{
			Status:  http.StatusInternalServerError,
			Message: defaultErrMsg})
	}
	return c.JSON(http.StatusOK, structs.DefaultResponse{
		Status:  http.StatusOK,
		Data:    users,
		Message: ""})
}

func (u UserController) Insert(c echo.Context) error {
	ur := repository.UserRepository{DB: u.DB}
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusInternalServerError, structs.ErrorReponse{
			Status:  http.StatusInternalServerError,
			Message: defaultErrMsg})
	}
	lastInsertID, err := ur.Insert(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, structs.ErrorReponse{
			Status:  http.StatusConflict,
			Message: defaultErrMsg})
	}
	if lastInsertID == 0 {
		return c.JSON(http.StatusInternalServerError, structs.ErrorReponse{
			Status:  http.StatusConflict,
			Message: "Este usuário já está cadastrado."})
	}
	return c.JSON(http.StatusOK, structs.DefaultResponse{
		Status:  http.StatusOK,
		Data:    map[string]int64{"lastInsertedID": lastInsertID},
		Message: ""})
}

func (u UserController) Login(c echo.Context) error {
	ur := repository.UserRepository{DB: u.DB}
	fmt.Println(ur)
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusInternalServerError, structs.ErrorReponse{
			Status:  http.StatusInternalServerError,
			Message: defaultErrMsg})
	}
	return c.JSON(http.StatusOK, structs.DefaultResponse{
		Status:  http.StatusInternalServerError,
		Data:    []int{1, 2},
		Message: ""})
}
