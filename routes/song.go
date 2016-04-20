package routes

import (
  "encoding/json"
  "net/http"
  "strconv"

  "github.com/gorilla/mux"

  "github.com/eecs394-s16/brown-api-golang/models"
)

func addSongRoutes(r *mux.Router) {
  // Get all songs
  r.HandleFunc("/songs", getSongsHandler).Methods("GET")

  // Create new song
  r.HandleFunc("/songs", createSongHandler).Methods("POST")

  // Upvote song
  r.HandleFunc("/songs/{song_id}/upvote", upvoteSongHandler).Methods("PUT")

  // Delete song
  r.HandleFunc("/songs/{song_id}", deleteSongHandler).Methods("DELETE")
}

func getSongsHandler(w http.ResponseWriter, req *http.Request) {
  // Get songs
  var songs []models.Song
  models.DB.Order("votes desc").Find(&songs)

  // Create response
  data := make(map[string]interface{})
  data["songs"] = songs

  setData(req, data)
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

  // Get songs sorted
  var songs []models.Song
  models.DB.Order("votes desc").Find(&songs)

  // Return songs sorted in response
  data := make(map[string]interface{})
  data["songs"] = songs
  json.NewEncoder(w).Encode(&data)
}

func upvoteSongHandler(w http.ResponseWriter, req *http.Request) {
  song_id, _ := strconv.Atoi(mux.Vars(req)["song_id"])

  song := models.SongFromID(song_id)

  // +1 to score
  song.Votes++

  // Save song
  models.DB.Save(&song)

  // Return song in response
  setData(req, song.GetData())
}

func deleteSongHandler(w http.ResponseWriter, req *http.Request) {
  song_id, _ := strconv.Atoi(mux.Vars(req)["song_id"])

  // Get song by id
  song := models.SongFromID(song_id)

  // Delete song
  models.DB.Delete(&song)

  data := make(map[string]interface{})
  data["deleted"] = song_id

  // Return response
  setData(req, data)
}
