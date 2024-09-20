package entity

import (
	"time"
)

type User struct {
	Id        int       `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	Username  string    `json:"username" db:"username"`
	CreatedAt time.Time `json:"created_at" db:"create_at"`
}

type AdminUser struct {
	Id       int    `json:"id" db:"id"`
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
}
