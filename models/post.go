package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID       int
	UserID   uint
	Caption  string
	ImageURL string
}
