package controllers

import (
	"air-server/form/api"
	"air-server/models/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct{}

// GET /users
func (pc UserController) Index(c echo.Context) error {
	var u repository.UserRepository
	p, err := u.GetAll()
	if err != nil {
		var apierr APIError
		apierr.Code = 404
		apierr.Message = err.Error()
		return c.JSON(http.StatusBadRequest, apierr)
	} else {
		return c.JSON(200, p)
	}
}

// POST /users
func (pc UserController) Create(c echo.Context) error {
	var u repository.UserRepository
	p, err := u.CreateModel(c)

	if err != nil {
		var apierr APIError
		apierr.Code = 400
		apierr.Message = err.Error()
		return c.JSON(http.StatusBadRequest, apierr)
	} else {
		return c.JSON(201, p)
	}
}

// POST /users/signup
func (pc UserController) Signup(c echo.Context) error {
	var u repository.UserRepository
	p, err := u.CreateModel(c)

	if err != nil {
		var apierr APIError
		apierr.Code = 400
		apierr.Message = err.Error()
		return c.JSON(http.StatusBadRequest, apierr)
	} else {
		return c.JSON(201, p)
	}
}

// POST /users/signin
func (pc UserController) Signin(c echo.Context) error {
	var u repository.UserRepository
	var user api.User
	// paramはget用だからpostはbindを使わないといけない
	if err := c.Bind(&user); err != nil {
		var apierr APIError
		apierr.Code = 400
		apierr.Message = err.Error()
		fmt.Println("run signin %v", c)

		return c.JSON(http.StatusBadRequest, apierr)
	}
	// idInt, _ := strconv.Atoi(user.ID)
	if user, err := u.SigninByIdAndPassword(user.ID, user.Password, c); err != nil {
		var apierr APIError
		apierr.Code = 400
		apierr.Message = err.Error()
		return c.JSON(http.StatusBadRequest, apierr)
	} else {
		return c.JSON(201, user)
	}
}

// Get /users/:id
func (pc UserController) Show(c echo.Context) error {
	id := c.Param("id")
	var u repository.UserRepository
	idInt, _ := strconv.Atoi(id)
	user, err := u.GetByID(idInt)

	if err != nil {
		var apierr APIError
		apierr.Code = 400
		apierr.Message = err.Error()
		return c.JSON(http.StatusBadRequest, apierr)
	} else {
		return c.JSON(200, user)
	}
}

// Put /users/:id
func (pc UserController) Update(c echo.Context) error {
	id := c.Param("id")
	var u repository.UserRepository
	idInt, _ := strconv.Atoi(id)
	p, err := u.UpdateByID(idInt, c)

	if err != nil {
		var apierr APIError
		apierr.Code = 404
		apierr.Message = err.Error()
		return c.JSON(http.StatusBadRequest, apierr)
	} else {
		return c.JSON(200, p)
	}
}

// DELETE /users/:id
func (pc UserController) Delete(c echo.Context) error {
	id := c.Param("id")
	var u repository.UserRepository
	if err := u.DeleteByID(id); err != nil {
		var apierr APIError
		apierr.Code = 403
		apierr.Message = err.Error()
		return c.JSON(http.StatusBadRequest, apierr)
	}

	return c.JSON(200, Success{success: "ID" + id + "のユーザーを削除しました"})
}
