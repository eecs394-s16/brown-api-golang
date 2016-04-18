package models

import (
  "github.com/jinzhu/gorm"
)

type Playlist struct {
  gorm.Model
  Title  string `json:"title"`
  Artist string `json:"artist"`
  Album  string `json:"album"`

  Votes  int    `json:"votes"`
}
