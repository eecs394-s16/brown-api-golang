package routes

import (
  "encoding/json"
  "net/http"

  "github.com/gorilla/mux"

  "github.com/eecs394-s16/brown-api-golang/models"

)

func addSongRoutes(r *mux.Router) {
  r.HandleFunc("/songs", getSongsHandler).Methods("GET")
  r.HandleFunc("/songs", createSongHandler).Methods("POST")
}

func getSongsHandler(w http.ResponseWriter, req *http.Request) {
  // Get songs
  var songs []models.Song
  models.DB.Find(&songs)

  // Create response
  w.Header().Set("Content-Type", "application/json")

  data := make(map[string]interface{})
  data["songs"] = songs
  json.NewEncoder(w).Encode(&data)
}

func createSongHandler(w http.ResponseWriter, req *http.Request) {

  // Get JSON request data
  decoder := json.NewDecoder(req.Body)
  var song models.Song
  err := decoder.Decode(&song)
  if err != nil {
    panic(err)
    return
  }

  // Save song
  models.DB.Create(&song)

  // Create response
  json.NewEncoder(w).Encode(&song)
}
