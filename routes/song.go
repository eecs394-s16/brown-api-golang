package routes

import (
  "encoding/json"
  "net/http"
  "fmt"

  "github.com/gorilla/mux"

  "github.com/eecs394-s16/brown-api-golang/models"

)

func addSongRoutes(r *mux.Router) {
  r.HandleFunc("/songs", getSongsHandler).Methods("GET")
  r.HandleFunc("/songs", createSongHandler).Methods("POST")

  r.HandleFunc("/songs/{song_id}/upvote", upvoteSongHandler).Methods("POST") // TODO
  r.HandleFunc("/songs/{song_id}", deleteSongHandler).Methods("DELETE") // TODO
}

// TODO
func deleteSongHandler(w http.ResponseWriter, req *http.Request) {
  song_id := mux.Vars(req)["song_id"]
  fmt.Println(song_id)

  // Get song by id

  // Delete song

  // Return response
   w.Write([]byte("delete songs2!\n"))
}

func getSongsHandler(w http.ResponseWriter, req *http.Request) {
  // Get songs
  var songs []models.Song
  models.DB.Order("votes desc").Find(&songs)

  // Create response
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

  // Initialize votes value
  song.Votes = 1

  // Save song
  models.DB.Create(&song)

  // Create response
  json.NewEncoder(w).Encode(&song)
}

func upvoteSongHandler(w http.ResponseWriter, req *http.Request) {
  song_id := mux.Vars(req)["song_id"]
  fmt.Println(song_id)

  var song models.Song
  models.DB.Where("ID = ?", song_id)

  // Check that song exists
  if models.DB.NewRecord(song) {
    fmt.Println("cannot find song")
  }
  fmt.Println("found song")

  // +1 to score

  // Save song

  // Return song in response
}
