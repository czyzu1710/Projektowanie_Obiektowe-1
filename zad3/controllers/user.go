package controllers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"zad3/dao"
	"zad3/model"
)

func GetUsers(c echo.Context) error {
	Users, err := dao.GetUsers()

	if err != nil {
		log.Println("Unable to fetch Users")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Users)

}

func GetUser(c echo.Context) error {
	User, err := dao.GetUser(c.Param("id"))

	if err != nil {
		log.Println("No such User")
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, User)

}

func CreateUser(c echo.Context) error {
	User := new(model.User)
	err := c.Bind(User)

	if err != nil {
		log.Println("Unable to create User.")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	dao.Save(User)
	return c.JSON(http.StatusOK, User)
}

func UpdateUser(c echo.Context) error {
	User := new(model.User)
	err := c.Bind(User)

	if err != nil {
		log.Println("Unable to update User.")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = dao.Update(User)

	if err != nil {
		log.Println("Error trying update User. No such User.")
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, User)
}

func DeleteUser(c echo.Context) error {
	err := dao.Delete(c.Param("id"))

	if err != nil {
		log.Println("Unable to delete User")
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, "")
}
