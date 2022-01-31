package controllers

import (
	"air-server/models/repository"
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
	idInt, _ := strconv.Atoi(id)
	if err := u.DeleteByID(idInt); err != nil {
		var apierr APIError
		apierr.Code = 403
		apierr.Message = err.Error()
		return c.JSON(http.StatusBadRequest, apierr)
	}

	return c.JSON(200, Success{success: "ID" + id + "のユーザーを削除しました"})
}
