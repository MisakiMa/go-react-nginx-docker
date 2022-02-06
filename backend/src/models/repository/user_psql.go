package repository

import (
	"air-server/crypto"
	"air-server/db"
	"air-server/models"

	"github.com/labstack/echo/v4"
)

type UserRepository struct{}

type User models.User

type UserProfile struct {
	Name string
	Id   int
}

// get all User
func (_ UserRepository) GetAll() ([]UserProfile, error) {
	db := db.GetDB()
	var u []UserProfile
	if err := db.Table("users").Select("name, id").Scan(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// create User model
func (_ UserRepository) CreateModel(c echo.Context) (User, error) {
	db := db.GetDB()
	var u User
	if err := c.Bind(&u); err != nil {
		return User{}, err
	}
	hash, err := crypto.PasswordEncrypt(u.Password)
	if err != nil {
		return User{}, err
	}
	u.Password = hash
	if err := db.Create(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

// get a User by ID
func (_ UserRepository) GetByID(id int) (models.User, error) {
	db := db.GetDB()
	var me models.User
	if err := db.Where("id = ?", id).First(&me).Error; err != nil {
		return me, err
	}
	var posts []models.Post
	db.Where("id = ?", id).First(&me)
	db.Model(&me).Related(&posts)
	me.Posts = posts

	return me, nil
}

// update a User
func (_ UserRepository) UpdateByID(id int, c echo.Context) (models.User, error) {
	db := db.GetDB()
	var u models.User
	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}
	if err := c.Bind(&u); err != nil {
		return u, err
	}
	u.ID = uint(id)
	db.Save(&u)

	return u, nil
}

// delete a User by ID
func (_ UserRepository) DeleteByID(id int) error {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}
