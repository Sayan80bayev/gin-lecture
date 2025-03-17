package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}
