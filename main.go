package main

import (
  _ "fmt"
  "github.com/eecs394-s16/brown-api-golang/models"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)



func main() {
  db := models.GetDB()
  
}
