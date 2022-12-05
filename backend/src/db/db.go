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

	// 仮ユーザーを追加する
	var u []models.User
	if err := db.Table("users").Select("user_name, id").Scan(&u).Error; err == nil || len(u) > 0 {
		return
	}
	user := models.User{
		UserID:   "12345",
		UserName: "aoki",
		Posts:    []models.Post{{ID: 1, Content: "comment1"}, {ID: 2, Content: "comment2"}},
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
