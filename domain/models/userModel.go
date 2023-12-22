package models

import "time"

type User struct {
	ID       int       `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	CreateAt string    `json:"create_at"`
	UpdateAt time.Time `json:"update_at" gorm:"autoUpdateTime"`
}

type CreateUser struct {
	ID       int       `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Password string    `json:"password"`
	CreateAt time.Time `json:"create_at" gorm:"autoCreateTime"`
}
