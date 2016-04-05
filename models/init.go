package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
  db, err := gorm.Open("postgres", "user=postgres dbname=jukebox password=bugzzues sslmode=disable")
  if err != nil {
    panic(err)
  }
  return db
}

func init() {
  var err interface{}
  DB, err = gorm.Open("postgres", "user=postgres dbname=jukebox password=bugzzues sslmode=disable")
  if err != nil {
    panic(err)
  }
}
