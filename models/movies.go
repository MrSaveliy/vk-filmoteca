package models

import "time"

type Movies struct {
	Id          int       `json:"-" db:"id"`
	Name        string    `json:"name" binding:"required"`
	Discription string    `json:"discription" binding:"required"`
	ReleaseDate time.Time `json:"release_date" binding:"required"`
	Rating      int       `json:"rating" binding:"required"`
}
