package routes

import (
  "net/http"

  "github.com/gorilla/mux"

  "github.com/codegangsta/negroni"
)

func GetRouter() *negroni.Negroni {
  n  := negroni.New()
  n.Use(negroni.HandlerFunc(configureResponseMiddleware))

  r  := mux.NewRouter()
  addSongRoutes(r)

  n.UseHandler(r)
  return n
}

func configureResponseMiddleware(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
  w.Header().Set("Content-Type", "application/json")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
  w.Header().Set("Access-Control-Allow-Headers",
    "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

    // Stop here for a Preflighted OPTIONS request.
  if req.Method == "OPTIONS" {
    return
  }
  next(w, req)
}
