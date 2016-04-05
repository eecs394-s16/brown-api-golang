package models

import (
  "github.com/jinzhu/gorm"
)

type Song struct {
  gorm.Model
  Title  string `json:"title"  gorm:"not null"`
  Artist string `json:"artist" gorm:"not null"`
  Album  string `json:"album"  gorm:"not null"`
}
