package main

import (
  "github.com/eecs394-s16/brown-api-golang/routes"
  _ "net/http/pprof"
  "net/http"
  "log"
)

func main() {
  go func() {
    log.Println(http.ListenAndServe("localhost:6060", nil))
  }()
  n  := routes.GetRouter()
  n.Run(":3000")
}
