package models

type User struct {
	ID       int    `json:"id" gorm:"AUTO_INCREMENT"`
	UserID   string `json:"userId" binding:"required"`
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
	Posts    []Post `json:"posts"`
}
