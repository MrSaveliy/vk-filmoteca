package models

type User struct {
	Id       int    `json:"-" db:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
	RoleId   int    `json:"role_id" binding:"required"`
}

type Role struct {
	Id          int    `json:"-" db:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
