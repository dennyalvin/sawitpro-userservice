package model

import "time"

type User struct {
	Id           int        `json:"id"`
	FullName     string     `json:"full_name"`
	Password     string     `json:"password"`
	Phone        string     `json:"phone"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	LoginSuccess uint       `json:"login_success"`
}
