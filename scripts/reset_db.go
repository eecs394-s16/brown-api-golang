package main

import (
  "fmt"
  "github.com/eecs394-s16/brown-api-golang/models"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
  // Add models here...
  _models := []interface{} {
    &models.Song{},
    &models.Playlist{}}

  fmt.Println("\nDropping tables...")
  for _, model := range _models {
    if !models.DB.HasTable(model) {
        fmt.Printf("* Skipping %T because it does not exist\n", model)
        continue
    }
    fmt.Printf("* Dropping %T...", model)
    models.DB.DropTable(model)
    fmt.Printf(" Dropped!\n")
  }

  fmt.Println("\nCreating tables...")
  for _, model := range _models {
    fmt.Printf("* Creating %T...", model)
    models.DB.CreateTable(model)
    fmt.Printf(" Created!\n")
  }

  fmt.Println("")
}
