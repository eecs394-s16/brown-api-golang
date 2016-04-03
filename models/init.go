package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDB() *gorm.DB {
  db, err := gorm.Open("postgres", "user=postgres dbname=jukebox password=bugzzues sslmode=disable")
  if err != nil {
    panic(err)
  }
  return db
}
