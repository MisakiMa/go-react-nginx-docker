package controllers

import (
	"air-server/form/api"
	"air-server/models"
	"air-server/models/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PostController struct{}

type APIError struct {
	Code    int
	Message string
}

type Success struct {
	success string
}

// GET /posts
func (pc PostController) Index(c echo.Context) error {
	var u repository.PostRepository
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

// POST /posts
func (pc PostController) Create(c echo.Context) error {
	var u repository.PostRepository
	in := api.Post{}
	if err := c.Bind(&in); err != nil {
		return err
	}
	in2 := &models.Post{
		ID:      in.ID,
		Content: in.Content,
		UserID:  in.UserID,
	}
	p, err := u.CreateModel(in2)
	if err != nil {
		var apierr APIError
		apierr.Code = 400
		return c.JSON(http.StatusBadRequest, apierr)
	} else {
		return c.JSON(201, p)
	}
}

// Get /posts/:id
func (pc PostController) Show(c echo.Context) error {
	id := c.Param("id")
	var p repository.PostRepository
	idInt, _ := strconv.Atoi(id)
	post, err := p.GetByID(idInt)

	if err != nil {
		var apierr APIError
		apierr.Code = 400
		return c.JSON(http.StatusBadRequest, apierr)
	} else {
		return c.JSON(200, post)
	}
}

// Put /posts/:id
func (pc PostController) Update(c echo.Context) error {
	id := c.Param("id")
	var u repository.PostRepository
	idInt, _ := strconv.Atoi(id)
	p, err := u.UpdateByID(idInt, c)

	if err != nil {
		var apierr APIError
		apierr.Code = 404
		return c.JSON(http.StatusBadRequest, apierr)
	} else {
		return c.JSON(200, p)
	}
}

// DELETE /posts/:id
func (pc PostController) Delete(c echo.Context) error {
	id := c.Param("id")
	var u repository.PostRepository
	idInt, _ := strconv.Atoi(id)
	if err := u.DeleteByID(idInt); err != nil {
		var apierr APIError
		apierr.Code = 403
		return c.JSON(http.StatusBadRequest, apierr)
	}

	return c.JSON(200, Success{success: "ID" + id + "の投稿を削除しました"})
}
