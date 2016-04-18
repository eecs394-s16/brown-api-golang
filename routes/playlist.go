package routes

import (
  _ "encoding/json"
  "net/http"
  "fmt"

  "github.com/gorilla/mux"
  "github.com/gorilla/context"

  "github.com/eecs394-s16/brown-api-golang/models"

)

func addPlaylistRoutes(r *mux.Router) {
  r.HandleFunc("/playlists/{playlist_id}", getSongsHandler).Methods("GET")
}

func getPlaylistHandler(w http.ResponseWriter, req *http.Request) {
  playlist_id := mux.Vars(req)["playlist_id"]

  // Get playlist
  var playlist models.Playlist
  models.DB.First(&playlist, playlist_id)

  // Check that song exists
  if models.DB.NewRecord(playlist) {
    fmt.Println("cannot find song")
    return
  }
  fmt.Println("found song")

  context.Set(req, "data", playlist)
}
