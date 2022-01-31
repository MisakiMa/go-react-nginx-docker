package server

import (
	"air-server/controllers"

	"github.com/labstack/echo/v4"
)

func Init() {
	e := router()
	e.Logger.Fatal(e.Start(":8080"))
}

func router() *echo.Echo {
	e := echo.New()

	userCtrl := controllers.UserController{}
	e.GET("/users", userCtrl.Index)
	e.POST("/users", userCtrl.Create)
	e.GET("/users/:id", userCtrl.Show)
	e.PUT("/users/:id", userCtrl.Update)
	e.DELETE("/users/:id", userCtrl.Delete)

	postCtrl := controllers.PostController{}
	e.GET("/posts", postCtrl.Index)
	e.POST("/posts", postCtrl.Create)
	e.GET("/posts/:id", postCtrl.Show)
	e.PUT("/posts/:id", postCtrl.Update)
	e.DELETE("/posts/:id", postCtrl.Delete)

	return e
}
