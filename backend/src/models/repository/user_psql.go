package repository

import (
	"air-server/crypto"
	"air-server/db"
	"air-server/models"
	"fmt"

	"github.com/labstack/echo/v4"
)

type UserRepository struct{}

type User models.User

type UserProfile struct {
	UserName string `json:"userName"`
	UserID   string `json:"userId"`
	ID       int    `json:"id"`
}

// get all User
func (UserRepository) GetAll() ([]UserProfile, error) {
	db := db.GetDB()
	var u []UserProfile
	if err := db.Table("users").Select("user_name, user_id, id").Scan(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// create User model
func (UserRepository) CreateModel(c echo.Context) (User, error) {
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
	// u.ID = uuid.NewString()
	if err := db.Create(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (UserRepository) SigninByIdAndPassword(id int, password string, e echo.Context) (models.User, error) {
	db := db.GetDB()
	fmt.Println("run signin %d %s", id, password)
	fmt.Println(id)
	fmt.Println(password)
	var me models.User
	if err := db.Table("users").Where("id = ?", id).First(&me).Error; err != nil {
		fmt.Println("not found user")
		return models.User{}, err
	}
	if err := crypto.CompareHashAndPassword(me.Password, password); err != nil {
		fmt.Println("not true compare hash and passowrd")
		return models.User{}, err
	}

	return me, nil
}

// get a User by ID
func (UserRepository) GetByID(id int) (models.User, error) {
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
func (UserRepository) UpdateByID(id int, c echo.Context) (models.User, error) {
	db := db.GetDB()
	var u models.User
	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}
	if err := c.Bind(&u); err != nil {
		return u, err
	}
	u.ID = id
	db.Save(&u)

	return u, nil
}

// delete a User by ID
func (UserRepository) DeleteByID(id string) error {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}
