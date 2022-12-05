package repository

import (
	"air-server/db"
	"air-server/form/api"
	"air-server/models"
	"fmt"

	"github.com/labstack/echo/v4"
)

type PostRepository struct{}

type Post api.Post

// get all Post
func (PostRepository) GetAll() ([]Post, error) {
	db := db.GetDB()
	var p []Post

	if err := db.Find(&p).Error; err != nil {
		return nil, err
	}

	return p, nil
}

// create Post model
func (PostRepository) CreateModel(p *models.Post) (*models.Post, error) {
	db := db.GetDB()
	if err := db.Create(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}

// get a Post by ID
func (PostRepository) GetByID(id int) (models.Post, error) {
	db := db.GetDB()
	var p models.Post
	if err := db.Where("id = ?", id).First(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}

// update a Post
func (PostRepository) UpdateByID(id int, c echo.Context) (api.Post, error) {
	db := db.GetDB()
	var p api.Post
	if err := db.Where("id = ?", id).First(&p).Error; err != nil {
		return p, err
	}
	userID := p.UserID
	if err := c.Bind(&p); err != nil {
		return p, err
	}
	fmt.Printf("%+V", p)
	p.ID = uint(id)
	p.UserID = userID
	db.Save(&p)

	return p, nil
}

// delete a Post by ID
func (PostRepository) DeleteByID(id int) error {
	db := db.GetDB()
	var p Post

	if err := db.Where("id = ?", id).Delete(&p).Error; err != nil {
		return err
	}

	return nil
}
