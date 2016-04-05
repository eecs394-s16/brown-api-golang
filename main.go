package main

import (
  "github.com/eecs394-s16/brown-api-golang/routes"
)

func main() {
  n  := routes.GetRouter()
  n.Run(":3000")
}
