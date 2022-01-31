package db

import (
	"air-server/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open("postgres", "host=db port=5432 user=hogehoge-api dbname=hogehoge-api password=hogehoge-api sslmode=disable")
	if err != nil {
		panic(err)
	}
	autoMigration()
	user := models.User{
		ID:    1,
		Name:  "aoki",
		Posts: []models.Post{{ID: 1, Content: "comment1"}, {ID: 2, Content: "comment2"}},
	}
	db.Create(&user)
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
}
