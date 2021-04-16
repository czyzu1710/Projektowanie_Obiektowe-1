package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"zad3/controllers"
	"zad3/database"
)

func main() {
	database.Init()
	db := database.GetInstance()
	db.Open()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users", controllers.GetUsers)
	e.POST("/users", controllers.CreateUser)
	e.GET("/users/:id", controllers.GetUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))
	db.Close()
}
