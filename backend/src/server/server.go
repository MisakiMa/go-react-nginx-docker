package server

import (
	"air-server/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() {
	e := router()
	e.Logger.Fatal(e.Start(":8080"))
}

func test(c echo.Context) error {
	return c.String(http.StatusOK, "Test!")
}

func router() *echo.Echo {
	e := echo.New()

	userCtrl := controllers.UserController{}
	e.GET("/users", userCtrl.Index)
	e.POST("/users", userCtrl.Create)
	e.GET("/users/:id", userCtrl.Show)
	e.PUT("/users/:id", userCtrl.Update)
	e.DELETE("/users/:id", userCtrl.Delete)
	e.POST("/users/signup", userCtrl.Signup)
	e.POST("/users/signin", userCtrl.Signin)

	postCtrl := controllers.PostController{}
	e.GET("/posts", postCtrl.Index)
	e.POST("/posts", postCtrl.Create)
	e.GET("/posts/:id", postCtrl.Show)
	e.PUT("/posts/:id", postCtrl.Update)
	e.DELETE("/posts/:id", postCtrl.Delete)

	e.GET("/test", test)

	return e
}
