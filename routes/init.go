package routes

import (
  "net/http"
  "encoding/json"

  "github.com/gorilla/mux"
  "github.com/gorilla/context"

  "github.com/codegangsta/negroni"
)

func GetRouter() *negroni.Negroni {
  n  := negroni.New()

  // Add configure response middleware
  n.Use(negroni.HandlerFunc(configureResponseMiddleware))
  n.Use(negroni.HandlerFunc(recoveryMiddleware()))

  r  := mux.NewRouter()
  r.KeepContext = true
  addSongRoutes(r)
  n.UseHandler(r)

  // Add handle response middleware
  n.Use(negroni.HandlerFunc(handleResponseMiddleware))
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

func handleResponseMiddleware(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
  data := getData(req)

  var status int
  if temp_status := getStatusCode(req); temp_status != nil {
    status = temp_status.(int)
  } else {
    status = 200
  }

  // Clear context
  context.Clear(req)

  // Set status code
  w.WriteHeader(status)

  // Respond with data
  json.NewEncoder(w).Encode(&data)

  next(w, req)
}

func recoveryMiddleware() negroni.HandlerFunc {
  return func(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
    defer func() {
      if err := recover(); err != nil {
        // Get status code
        var status int
        if temp_status := getStatusCode(req); temp_status != nil {
          status = temp_status.(int)
        } else {
          status = 500
        }
        // Set status code
        w.WriteHeader(status)

        // Write error back
        data := make(map[string]interface{})
        data["status_code"] = status
        data["error"] = err
        json.NewEncoder(w).Encode(&data)
      }
    }()

    next(w, req)
  }
}

// Get/Set Context Variables
func setStatusCode(req *http.Request, status int) {
  context.Set(req, 0, status)
}

func getStatusCode(req *http.Request) interface{} {
  return context.Get(req, 0)
}

func setData(req *http.Request, data interface{}) {
  context.Set(req, 1, data)
}

func getData(req *http.Request) interface{} {
  return context.Get(req, 1)
}
