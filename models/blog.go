package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Viewer      int    `json:"viewer"`
}
