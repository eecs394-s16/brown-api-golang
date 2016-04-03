package main

import (
  "fmt"
  "github.com/eecs394-s16/brown-api-golang/models"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
  db := models.GetDB()

  // Add models here...
  _models := []interface{} {
    &models.Song{}}

  fmt.Println("\nDropping tables...")
  for _, model := range _models {
    if !db.HasTable(model) {
        fmt.Printf("* Skipping %T because it does not exist\n", model)
        continue
    }
    fmt.Printf("* Dropping %T...", model)
    db.DropTable(model)
    fmt.Printf(" Dropped!\n")
  }

  fmt.Println("\nCreating tables...")
  for _, model := range _models {
    fmt.Printf("* Creating %T...", model)
    db.CreateTable(model)
    fmt.Printf(" Created!\n")
  }

  fmt.Println("")
}
