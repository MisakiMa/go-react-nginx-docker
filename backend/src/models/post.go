package models

type Post struct {
	ID      uint   `json:"id" binding:"required"`
	Content string `json:"content" binding:"required"`
	User    User   `json:"-" binding:"required"`
	UserID  uint   `gorm:"not null" json:"userId"`
}
