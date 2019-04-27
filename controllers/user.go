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

func getErrorReponse(c echo.Context, message string) error {
	return c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: message})
}

func getSuccessResponse(c echo.Context, data interface{}, message string) error {
	return c.JSON(http.StatusInternalServerError, structs.DefaultResponse{
		Status:  http.StatusOK,
		Data:    data,
		Message: message})
}

func (u UserController) FindAll(c echo.Context) error {
	ur := repository.UserRepository{DB: u.DB}
	users, err := ur.FindAll()
	if err != nil {
		return getErrorReponse(c, defaultErrMsg)
	}
	return getSuccessResponse(c, users, "")
}

func (u UserController) Insert(c echo.Context) error {
	ur := repository.UserRepository{DB: u.DB}
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return getErrorReponse(c, defaultErrMsg)
	}
	lastInsertID, err := ur.Insert(user)
	if err != nil {
		return getErrorReponse(c, defaultErrMsg)
	}
	if lastInsertID == 0 {
		return getErrorReponse(c, "Este usuário já está cadastrado.")
	}
	return getSuccessResponse(c, map[string]int64{"lastInsertedID": lastInsertID}, "")
}

func (u UserController) Login(c echo.Context) error {
	ur := repository.UserRepository{DB: u.DB}
	fmt.Println(ur)
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return getErrorReponse(c, defaultErrMsg)
	}

	userResult, err := ur.Login(user.Email, user.Password)
	if err != 0 {
		if err == 1 {
			return getErrorReponse(c, defaultErrMsg)
		} else if err == 2 {
			return getErrorReponse(c, "Usuário não encontrado.")
		} else if err == 3 {
			return getErrorReponse(c, defaultErrMsg)
		} else if err == 4 {
			return getErrorReponse(c, "Senha incorreta.")
		}
	}

	return getSuccessResponse(c, userResult, "")
}
