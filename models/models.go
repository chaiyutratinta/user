package models

type User struct {
	ID       string `gorm:"column:user_id"`
	UserName string `gorm:"column:user_name"`
	Password string `gorm:"column:password"`
	Email    string `gorm:"column:email"`
}

type Body struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `josn:"email"`
}

type Token string
