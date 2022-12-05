package api

type User struct {
	ID       string `json:"id" binding:"required"`
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserPosts struct {
	ID       string `json:"id" binding:"required"`
	UserName string `json:"userName" binding:"required"`
	Posts    []Post `json:"posts"`
}
